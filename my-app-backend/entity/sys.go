package entity

import (
	"github.com/SuanCaiYv/my-app-backend/util"
	"time"
)

type SysUser struct {
	Id          string    `bson:"_id" json:"user_id"`
	Username    string    `bson:"username" json:"username"`
	Credential  string    `bson:"credential" json:"-"`
	Salt        string    `bson:"salt" json:"-"`
	Role        string    `bson:"role" json:"role"`
	Info        UserInfo  `bson:"info" json:"info"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type SysRole struct {
	Id          string    `bson:"_id" json:"-"`
	Name        string    `bson:"name" json:"name"`
	Desc        string    `bson:"desc" json:"desc"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

func DefaultSysUser() *SysUser {
	return &SysUser{
		Username:   "default-username" + util.GenerateUUID(),
		Credential: "default-credential" + util.GenerateUUID(),
		Salt:       util.GenerateUUID(),
		Role:       "default-role",
		Info: UserInfo{
			Avatar:    "http://127.0.0.1:8190/v1/static/a/default-avatar.png",
			Nickname:  "default-nickname-" + util.GenerateUUID(),
			Email:     "default-email",
			Phone:     "15651731700",
			WeChat:    "d2508826394",
			QQ:        "2508826394",
			Github:    "https://github.com/SuanCaiYv",
			Location:  "HangZhou-China",
			Signature: "Golang is best!",
		},
	}
}

const (
	RoleOwner  = "owner"
	RoleReader = "reader"
)
