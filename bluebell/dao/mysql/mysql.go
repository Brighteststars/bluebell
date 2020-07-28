package mysql

import "C"
import (
	"bluebell/models"
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义一个全局对象db
var db *gorm.DB

// 定义一个初始化数据库的函数
func Init() (err error) {
	mysqlConfig := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"))

	//db, err = gorm.Open("mysql",
	//	"dmfs9000:dmfs9000@(127.0.0.1:3306)/dmfs9000?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql",
		mysqlConfig)
	if err != nil {
		fmt.Printf("gorm.Open() connect mysql failed, err: %v\n", err)
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 创建表 自动迁移 (结构体与数据库表对应)
	//db.AutoMigrate(&User{}, &Login{}, &Right{}, &Role{}, &RoleRight{})
	db.AutoMigrate(&models.User{})

	return
}

func Close() {
	db.Close()
}
