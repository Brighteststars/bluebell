package mysql

import (
	"bluebell/models"
	"fmt"

	"github.com/jinzhu/gorm"

	"go.uber.org/zap"
)

func Login(user *models.User) (err error) {
	// origin登录密码
	lgPwd := user.Password

	err = db.Select("password,user_id,username").Where("username = ?", user.Username).First(user).Error
	if gorm.IsRecordNotFoundError(err) {
		// 数据库没有找到数据
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("用户登录sql执行失败", zap.Error(err))
		return
	}
	fmt.Printf("user: %#v\n", *user)

	oPassword := encryptPassword(lgPwd)
	if oPassword != user.Password {
		return ErrorInvalidPassword
	}
	return
}
