package config

import (
	"encoding/json"
	"flag"
	"github.com/SuanCaiYv/my-app-backend/util"
	"os"
)

func init() {
	configPath := flag.String("config", "", "config file path")
	flag.StringVar(configPath, "c", "", "config file path")
	flag.Parse()
	if *configPath == "" {
		util.JustPanic("config path is empty")
	}
	confFile, err := os.OpenFile(*configPath, os.O_RDONLY, os.ModePerm)
	util.JustPanic(err)
	fileInfo, err := confFile.Stat()
	util.JustPanic(err)
	bytes := make([]byte, fileInfo.Size(), fileInfo.Size())
	_, err = confFile.Read(bytes)
	util.JustPanic(err)
	configObject := make(map[string]interface{})
	err = json.Unmarshal(bytes, &configObject)
	util.JustPanic(err)

	database := configObject["database"].(map[string]interface{})
	tmpRoles := configObject["roles"].([]interface{})
	roles := make([]Role, 0, len(tmpRoles))
	for _, val := range tmpRoles {
		role := val.(map[string]interface{})
		roles = append(roles, Role{
			Name: role["name"].(string),
			Desc: role["desc"].(string),
		})
	}
	redis := configObject["redis"].(map[string]interface{})
	tmpAccounts := configObject["accounts"].([]interface{})
	accounts := make([]Account, 0, len(tmpAccounts))
	accountSet := make(Set)
	for _, val := range tmpAccounts {
		tmpAccount := val.(map[string]interface{})
		accounts = append(accounts, Account{
			Username:   tmpAccount["username"].(string),
			Credential: tmpAccount["credential"].(string),
			VerCode:    tmpAccount["ver_code"].(string),
		})
		accountSet[tmpAccount["username"].(string)] = struct{}{}
	}
	email := configObject["email"].(map[string]interface{})
	config = &Configuration{
		Owner: configObject["owner"].(string),
		DatabaseConfig: &DatabaseConfig{
			Url:      database["url"].(string),
			Port:     int(database["port"].(float64)),
			DB:       database["db"].(string),
			GridFSDB: database["grid_fs_db"].(string),
			User:     database["user"].(string),
			Password: database["password"].(string),
		},
		RedisConfig: &RedisConfig{
			Url:      redis["url"].(string),
			Port:     int(redis["port"].(float64)),
			DB:       int(redis["db"].(float64)),
			User:     redis["user"].(string),
			Password: redis["password"].(string),
		},
		Roles:      roles,
		Accounts:   accounts,
		AccountSet: accountSet,
		Email: &Email{
			Host:       email["host"].(string),
			Port:       int(email["port"].(float64)),
			Sender:     email["sender"].(string),
			Credential: email["credential"].(string),
		},
	}
}

type Set map[string]struct{}

type Configuration struct {
	Owner          string
	DatabaseConfig *DatabaseConfig
	RedisConfig    *RedisConfig
	Roles          []Role
	Accounts       []Account
	AccountSet     Set
	Email          *Email
}

type DatabaseConfig struct {
	Url      string
	Port     int
	DB       string
	GridFSDB string
	User     string
	Password string
}

type RedisConfig struct {
	Url      string
	Port     int
	DB       int
	User     string
	Password string
}

type Role struct {
	Name string
	Desc string
}

type Account struct {
	Username   string
	Credential string
	VerCode    string
}

type Email struct {
	Host       string
	Port       int
	Sender     string
	Credential string
}

var config *Configuration

func ApplicationConfiguration() *Configuration {
	return config
}
