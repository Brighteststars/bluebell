package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
)

func Login(p *models.ParamsLogin) (token string, err error) {
	// 用户登录处理
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.Login(user); err != nil {
		// 登录成功

		return "", err
	}

	// 生成JWT

	return jwt.GenToken(user.UserId, user.Username)
}
