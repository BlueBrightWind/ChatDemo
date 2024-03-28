package mysql

import (
	"ChatDemo/global"
	"ChatDemo/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	//连接数据库
	config := global.Config.MySQLConfig
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.Name, config.Password, config.Host, config.Port, config.DataBase, config.Charset)
	db, _ := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	return db
}

func InitMySQL() {
	DB.AutoMigrate(model.User{}, model.Community{}, model.Contact{})
}

var DB = NewConnection()
