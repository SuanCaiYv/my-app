package service

import (
	"crypto/md5"
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/auth"
	"github.com/SuanCaiYv/my-app-backend/config"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/entity"
	"github.com/SuanCaiYv/my-app-backend/entity/resp"
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type UserApi interface {
	SignUp(context *gin.Context)

	Login(context *gin.Context)

	Logout(context *gin.Context)

	GetUserInfo(context *gin.Context)

	UpdateUserInfo(context *gin.Context)
}

type UserApiHandler struct {
	sysUserDao db.SysUserDao
	gridFsDao  db.GridFSDao
	redisOps   nosql.RedisOps
	logger     *logrus.Logger
}

type sign struct {
	Username   string `json:"username"`
	Credential string `json:"credential"`
	VerCode    string `json:"verCode"`
}

func NewUserApiHandler() *UserApiHandler {
	return &UserApiHandler{
		sysUserDao: db.NewSysUserDaoService(),
		gridFsDao:  db.NewGridFSDaoService(),
		redisOps:   nosql.NewRedisClient(),
		logger:     util.NewLogger(),
	}
}

func (u *UserApiHandler) SignUp(context *gin.Context) {
	sign := &sign{}
	err := context.BindJSON(sign)
	if err != nil {
		u.logger.Errorf("参数解析失败: %v", err)
		context.JSON(200, resp.NewBadRequest("参数解析失败"))
		return
	}
	sysUser, err := u.sysUserDao.SelectByUsername(sign.Username)
	if err != nil {
		u.logger.Errorf("无法读取SysUser数据表: %v", err)
		context.JSON(200, resp.NewIntervalError("无法读取用户表"))
		return
	}
	verCodeCache := ""
	_ = u.redisOps.Get("ver_code_"+sign.Username, &verCodeCache)
	if verCodeCache != sign.VerCode {
		u.logger.Infof("验证码错误: %s", sign.Username)
		context.JSON(200, resp.NewBadRequest("验证码错误"))
		return
	}
	salt := strings.ReplaceAll(util.GenerateUUID(), "-", "")[:6]
	secretPwd := fmt.Sprintf("%x", md5.Sum([]byte(sign.Credential+"_"+salt)))
	if sysUser != nil {
		sysUser.Credential = secretPwd
		sysUser.Salt = salt
		err := u.sysUserDao.Update(sysUser)
		if err != nil {
			u.logger.Errorf("更新SysUser失败: %s", sign.Username)
			context.JSON(200, resp.NewIntervalError("更新用户数据表失败"))
			return
		}
	} else {
		sysUser = entity.DefaultSysUser()
		sysUser.Username = sign.Username
		sysUser.Credential = secretPwd
		sysUser.Salt = salt
		if sign.Username == config.ApplicationConfiguration().Owner {
			sysUser.Role = entity.RoleOwner
		} else {
			sysUser.Role = entity.RoleReader
		}
		err := u.sysUserDao.Insert(sysUser)
		if err != nil {
			u.logger.Errorf("插入SysUser失败: %s", sign.Username)
			context.JSON(200, resp.NewIntervalError("插入用户数据表失败"))
			return
		}
	}
	context.JSON(200, resp.NewOk(&struct{}{}))
}

func (u *UserApiHandler) Login(context *gin.Context) {
	sign := &sign{}
	err := context.BindJSON(sign)
	if err != nil {
		u.logger.Errorf("参数解析失败: %v", err)
		context.JSON(200, resp.NewBadRequest("参数解析失败"))
		return
	}
	sysUser, err := u.sysUserDao.SelectByUsername(sign.Username)
	if err != nil {
		u.logger.Errorf("无法读取SysUser数据表: %v", err)
		context.JSON(200, resp.NewIntervalError("无法读取用户数据表"))
		return
	}
	if sysUser == nil {
		u.logger.Infof("用户不存在: %s", sign.Username)
		context.JSON(200, resp.NewBadRequest("用户不存在"))
		return
	}
	secretPwd := fmt.Sprintf("%x", md5.Sum([]byte(sign.Credential+"_"+sysUser.Salt)))
	if secretPwd != sysUser.Credential {
		u.logger.Infof("密码错误: %s", sign.Username)
		context.JSON(200, resp.NewBadRequest("密码错误"))
		return
	}
	refreshToken, err := auth.SignRefreshToken(sysUser.Username)
	if err != nil {
		u.logger.Errorf("生成RefreshToken失败: %s", sign.Username)
		context.JSON(200, resp.NewIntervalError("生成令牌失败"))
		return
	}
	accessToken, err := auth.SignAccessToken(sysUser.Username, sysUser.Role)
	if err != nil {
		u.logger.Errorf("AccessToken: %s", sign.Username)
		context.JSON(200, resp.NewIntervalError("生成令牌失败"))
		return
	}
	context.JSON(200, resp.NewOk(&struct {
		RefreshToken string `json:"refreshToken"`
		AccessToken  string `json:"accessToken"`
	}{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}))
}

func (u *UserApiHandler) Logout(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) GetUserInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) UpdateUserInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
