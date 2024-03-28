package global

type MYSQL_MODEL struct {
	Name     string
	Password string
	Host     string
	Port     int
	DataBase string
	Charset  string
}

type REDIS_MODEL struct {
	Host        string
	Port        int
	Password    string
	Db          int
	PoolSize    int
	MinIdleConn int
}

type SOCKET_MODEL struct {
	HeartBeatInterval int
	HeartbeatMaxTime  int
}

type SERVER_MODEL struct{}

type CONFIG struct {
	MySQLConfig   MYSQL_MODEL
	RedisConfig   REDIS_MODEL
	SocketConfig  SOCKET_MODEL
	ServiceConfig SERVER_MODEL
}
