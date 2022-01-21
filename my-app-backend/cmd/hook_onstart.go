package main

import (
	"crypto/md5"
	"fmt"
	config2 "github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/nosql"
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
	gridFsDao := db.NewGridFSDaoService()
	if !gridFsDao.ExistFile("default-avatar.png") {
		defaultAvatarPath, err := filepath.Abs("static/default-avatar.png")
		util.JustPanic(err)
		defaultAvatar, err := os.OpenFile(defaultAvatarPath, os.O_RDONLY, os.ModePerm)
		util.JustPanic(err)
		data, err := ioutil.ReadAll(defaultAvatar)
		util.JustPanic(err)
		err = gridFsDao.UploadFile(data, "default-avatar.png", make(map[string]interface{}))
		util.JustPanic(err)
	}
	_ = mime.AddExtensionType(".md", "text/x-markdown")
}
