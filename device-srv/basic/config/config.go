package config

type AppConfig struct {
	Mysql         `json:"mysql"`
	Redis         `json:"redis"`
	MongoDb       `json:"mongoDb"`
	Elasticsearch `json:"elasticsearch"`
	RabbitMq      `json:"rabbitMq"`
}

type RabbitMq struct {
	Host     string
	Port     int
	User     string
	Password string
	Vhost    string
}

type Elasticsearch struct {
	Host string
	Port int
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
