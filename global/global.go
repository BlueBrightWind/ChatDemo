package global

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() CONFIG {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	var mysql_config = MYSQL_MODEL{
		Name:     viper.GetString("MySQL.Name"),
		Password: viper.GetString("MySQL.Password"),
		Host:     viper.GetString("MySQL.Host"),
		Port:     viper.GetInt("MySQL.Port"),
		DataBase: viper.GetString("MySQL.Database"),
		Charset:  viper.GetString("MySQL.Charset"),
	}
	var redis_config = REDIS_MODEL{
		Host:        viper.GetString("Redis.Host"),
		Port:        viper.GetInt("Redis.Port"),
		Password:    viper.GetString("Redis.Password"),
		Db:          viper.GetInt("Redis.DB"),
		PoolSize:    viper.GetInt("Redis.Poolsize"),
		MinIdleConn: viper.GetInt("Redis.Minidleconn"),
	}
	var socket_config = SOCKET_MODEL{
		HeartBeatInterval: viper.GetInt("socket.HeartBeatInterval"),
		HeartbeatMaxTime:  viper.GetInt("socket.HeartbeatMaxTime"),
	}
	var server_config = SERVER_MODEL{}
	config := CONFIG{
		MySQLConfig:   mysql_config,
		RedisConfig:   redis_config,
		SocketConfig:  socket_config,
		ServiceConfig: server_config,
	}
	return config
}

var Config = InitConfig()
