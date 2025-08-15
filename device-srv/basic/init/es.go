package init

import (
	"device-srv/basic/config"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func InitEs() {
	addr := fmt.Sprintf("http://%s:%d", config.AppConf.Elasticsearch.Host, config.AppConf.Elasticsearch.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			addr,
		},
		// ...
	}
	var err error
	config.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("es init success")
}
