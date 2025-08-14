package config

type AppConfig struct {
	Mysql   `json:"mysql"`
	Redis   `json:"redis"`
	MongoDb `json:"mongoDb"`
}

type MongoDb struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int
	User     string
	Password string
	Database int
}
