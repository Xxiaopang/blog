package config

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type server struct {
	Mode string
	Port string
}
type redis struct {
	Addr     string
	Password string
	DB       int
}

type mysqlDB struct {
	User   string
	Pwd    string
	Host   string
	Port   string
	Db     string
	Option string
}

var JWTKey string
var Header string

var App = &server{}
var MysqlDB = &mysqlDB{}
var Redis = &redis{}

func init() {
	wordDir, err := os.Getwd()
	if err != nil {
		return
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "./config")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("Read In Config error: " + err.Error())
		panic("Read In Config error" + err.Error())
		return
	}

}

func InitConfig() {
	JWTKey = viper.GetString("jwt.key")
	Header = viper.GetString("jwt.header")

	App.Mode = viper.GetString("server.mode")
	App.Port = viper.GetString("server.port")

	MysqlDB.User = viper.GetString("database.user")
	MysqlDB.Pwd = viper.GetString("database.pwd")
	MysqlDB.Host = viper.GetString("database.host")
	MysqlDB.Port = viper.GetString("database.port")
	MysqlDB.Db = viper.GetString("database.db")
	MysqlDB.Option = viper.GetString("database.option")

	Redis.Addr = viper.GetString("redis.addr")
	Redis.Password = viper.GetString("redis.pwd")
	Redis.DB = viper.GetInt("redis.db")
}
