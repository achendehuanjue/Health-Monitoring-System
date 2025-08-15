package init

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"pex-srv/basic/config"
	"time"
)

func InitEmqx() {
	config.Opts = mqtt.NewClientOptions()
	config.Opts.AddBroker("tcp://broker.emqx.io:1883")

	config.Opts.SetCleanSession(true)
	config.Opts.SetAutoReconnect(true)
	config.Opts.SetMaxReconnectInterval(5 * time.Second)

}
