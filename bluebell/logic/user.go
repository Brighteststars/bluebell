package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamsSignUp) (err error) {
	// 1.判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return
	}

	// 2.生成uId
	uid := snowflake.GenID()
	u := &models.User{
		UserId:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	// 3.保存到数据库
	err = mysql.InsertUser(u)

	return
}
