package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"go.uber.org/zap"
)

const secret = "one piece"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func CheckUserExist(username string) (err error) {

	var count int
	// 查不到数据 err!= nil
	err = db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		zap.L().Error("查询用户名称是否存在sql执行失败", zap.Error(err))
		return
	}

	if count != 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// password md5加密
	pwd := encryptPassword(user.Password)
	user.Password = pwd
	// 创建记录
	if err = db.Create(user).Error; err != nil {
		return
	}

	return
}

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
