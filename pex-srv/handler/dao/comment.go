package dao

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"pex-srv/basic/config"
	"pex-srv/handler/pkg"
	"time"
)

// 体温
func TemperatureEmqxClient() {
	clientID := fmt.Sprintf("mqtt-consumer-Temperature-%d", time.Now().UnixNano())

	config.Opts.SetClientID(clientID)
	config.Opts.SetDefaultPublishHandler(pkg.TemperatureMessageHandler)
	config.Opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("重新连接到MQTT服务器")
	})

	pkg.MqttClient(config.Temperature)

}

// 血糖
func BloodGlucoseEmqxClient() {
	clientID := fmt.Sprintf("mqtt-consumer-BloodGlucose-%d", time.Now().UnixNano())

	config.Opts.SetClientID(clientID)
	config.Opts.SetDefaultPublishHandler(pkg.BloodGlucoseMessageHandler)
	config.Opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("重新连接到MQTT服务器")
	})

	pkg.MqttClient(config.BloodGlucose)
}

// 血压
func BloodPressureEmqxClient() {
	clientID := fmt.Sprintf("mqtt-consumer-BloodPressure-%d", time.Now().UnixNano())

	config.Opts.SetClientID(clientID)
	config.Opts.SetDefaultPublishHandler(pkg.BloodPressureMessageHandler)
	config.Opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("重新连接到MQTT服务器")
	})

	pkg.MqttClient(config.BloodPressure)
}

// 血氧
func BloodOxygenEmqxClient() {
	clientID := fmt.Sprintf("mqtt-consumer-BloodOxygen-%d", time.Now().UnixNano())

	config.Opts.SetClientID(clientID)
	config.Opts.SetDefaultPublishHandler(pkg.BloodOxygenMessageHandler)
	config.Opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("重新连接到MQTT服务器")
	})

	pkg.MqttClient(config.BloodOxygen)
}
