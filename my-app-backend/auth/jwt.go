package auth

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/nosql"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var redisOps = nosql.NewRedisClient()
var logger = util.NewLogger()

func SignAccessToken(username, role string) (refreshToken string, err error) {
	secretKey := ""
	err = redisOps.Get(username, &secretKey)
	if err != nil && err != redis.Nil {
		logger.Errorf("获取JwtId异常: %v", err)
		return "", err
	}
	if secretKey == AccountLocked {
		logger.Errorf("账户已锁定: %s", username)
		return "", fmt.Errorf("账户已锁定: %s", username)
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "CWB",
		Subject:   "AccessToken",
		Audience:  []string{0: username, 1: role},
		ExpiresAt: jwt.NewNumericDate(time.UnixMilli(time.Now().UnixMilli() + 2*time.Hour.Milliseconds())),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        username,
	}).SignedString([]byte(secretKey))
	if err != nil {
		logger.Error("签发访问令牌失败")
		return "", err
	}
	return refreshToken, nil
}

func SignRefreshToken(username string) (accessToken string, err error) {
	secretKey := ""
	err = redisOps.Get(username, &secretKey)
	if err != nil && err != redis.Nil {
		logger.Infof("获取JwtId异常: %v", err)
		return "", err
	}
	if secretKey == AccountLocked {
		logger.Errorf("账户已锁定: %s", username)
		return "", fmt.Errorf("账户已锁定: %s", username)
	}
	secretKey = util.GenerateUUID()
	err = redisOps.SetExp(username, secretKey, 7*24*time.Hour)
	if err != nil {
		logger.Errorf("设置JwtId异常: %v", err)
		return "", err
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "CWB",
		Subject:   "RefreshToken",
		Audience:  []string{0: "RefreshToken"},
		ExpiresAt: jwt.NewNumericDate(time.UnixMilli(time.Now().UnixMilli() + 7*24*time.Hour.Milliseconds())),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        username,
	}).SignedString([]byte(secretKey))
	if err != nil {
		logger.Error("签发刷新令牌失败")
		return "", err
	}
	return accessToken, err
}

func ValidToken(token string) (username string, role string, err error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			secret := ""
			username = claims.ID
			err = redisOps.Get(username, &secret)
			if err != nil {
				logger.Warn(err)
			}
			fmt.Println(secret)
			return []byte(secret), nil
		}
		return []byte("Failed!"), nil
	})
	if err != nil {
		logger.Errorf("解析Token失败: %v", err)
		return "", "", err
	}
	if claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims); ok {
		username = claims.ID
		role = ""
		if len(claims.Audience) > 1 {
			role = claims.Audience[1]
		}
		return username, role, nil
	} else {
		return "", "", err
	}
}

const (
	AccountLocked = "__CwB__LOCKED__cWb__"
)
