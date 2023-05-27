package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"project/config"
	"project/model"
)

var DB *gorm.DB

func InitDB() {
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.MysqlDB.User,
		config.MysqlDB.Pwd,
		config.MysqlDB.Host,
		config.MysqlDB.Port,
		config.MysqlDB.Db,
		config.MysqlDB.Option,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db open error: " + err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.Comment{})

	DB = db
}

func Close() {
	db, _ := DB.DB()

	defer db.Close()
}
