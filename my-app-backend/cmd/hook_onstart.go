package main

import (
	"crypto/md5"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/service"
	"github.com/SuanCaiYv/my-app-backend/util"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func BeforeStart() {
	config := config2.ApplicationConfiguration()
	// 添加角色
	sysRoleDao := db.NewSysRoleDaoService()
	for _, val := range config.Roles {
		result, err := sysRoleDao.SelectByName(val.Name)
		util.JustPanic(err)
		if result == nil {
			role := &entity.SysRole{
				Name: val.Name,
				Desc: val.Desc,
			}
			err := sysRoleDao.Insert(role)
			util.JustPanic(err)
		}
	}
	// 设置验证码，添加用户
	sysUserDao := db.NewSysUserDaoService()
	redisOps := nosql.NewRedisClient()
	for _, val := range config.Accounts {
		err := redisOps.SetExp("ver_code_"+val.Username, val.VerCode, 7*24*time.Hour)
		util.JustPanic(err)
		tmpSysUser, err := sysUserDao.SelectByUsername(val.Username)
		util.JustPanic(err)
		if tmpSysUser != nil {
			continue
		}
		sysUser := entity.DefaultSysUser()
		sysUser.Username = val.Username
		sysUser.Salt = strings.ReplaceAll(util.GenerateUUID(), "-", "")[:6]
		sysUser.Credential = fmt.Sprintf("%x", md5.Sum([]byte(val.Credential+"_"+sysUser.Salt)))
		if config.Owner == val.Username {
			sysUser.Role = entity.RoleOwner
		} else {
			sysUser.Role = entity.RoleReader
		}
		sysUser.Info.Email = val.Username
		sysUser.Info.Nickname = val.Username
		err = sysUserDao.Insert(sysUser)
		util.JustPanic(err)
	}
	// 设置默认头像
	gridFsDao := db.NewGridFSDaoService()
	filename := "default-avatar.png"
	if !gridFsDao.ExistFile(filename) {
		defaultAvatarPath, err := filepath.Abs("static/" + filename)
		util.JustPanic(err)
		defaultAvatar, err := os.OpenFile(defaultAvatarPath, os.O_RDONLY, os.ModePerm)
		util.JustPanic(err)
		fileInfo, err := defaultAvatar.Stat()
		util.JustPanic(err)
		metaMap := make(map[string]interface{})
		metaMap["upload_user"] = config2.ApplicationConfiguration().Owner
		metaMap["origin_name"] = fileInfo.Name()
		metaMap["archive"] = "avatar"
		data, err := ioutil.ReadAll(defaultAvatar)
		util.JustPanic(err)
		err = gridFsDao.UploadFile(data, filename, metaMap)
		util.JustPanic(err)
	}
	// 设置文件后缀识别
	_ = mime.AddExtensionType(".md", "text/x-markdown")
	// 启动文章清除器
	clearFunc := func(params service.Params) {
		articleDao := db.NewArticleDaoService()
		sysUserDao := db.NewSysUserDaoService()
		user, err := sysUserDao.SelectByUsername(config2.ApplicationConfiguration().Owner)
		util.JustPanic(err)
		articles, err := articleDao.ListByAuthor0(user.Id, entity.VisibilityDraft, true)
		util.JustPanic(err)
		curr := time.Now()
		for i := range articles {
			if articles[i].UpdatedTime.Before(curr.Add(-2*time.Minute)) && len(articles[i].Name) == 0 && len(articles[i].Content) == 0 {
				err = articleDao.Delete(articles[i].Id)
				util.JustPanic(err)
			}
		}
		service.Add("clearFunc", make(service.Params), time.Now().Add(2*time.Minute))
	}
	service.AddFunction("clearFunc", clearFunc)
	service.Add("clearFunc", make(service.Params), time.Now())
}
