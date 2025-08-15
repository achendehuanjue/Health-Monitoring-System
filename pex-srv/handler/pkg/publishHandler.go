package pkg

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"pex-srv/basic/config"
	"pex-srv/handler/model"
	"time"
)

func MqttClient(str string) {
	client := mqtt.NewClient(config.Opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("连接失败: %v", token.Error()))
	}

	token := client.Subscribe("health/+/"+str, 1, nil)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("订阅主题失败: %v\n", token.Error())
	}
}

// 体温消息处理函数
var TemperatureMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()

	var rawData model.RawData
	json.Unmarshal(payload, &rawData)

	rawData.Metric = config.Temperature
	rawData.Unit = "℃"
	rawData.Timestamp = time.Now().Unix()
	fmt.Println(rawData)
}

// 血糖消息处理函数
var BloodGlucoseMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()

	var rawData model.RawData
	json.Unmarshal(payload, &rawData)

	rawData.Metric = config.BloodGlucose
	rawData.Unit = "mg/dL"
	rawData.Timestamp = time.Now().Unix()
	fmt.Println(rawData)
}

// 血压消息处理函数
var BloodPressureMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()

	var rawData model.RawData
	json.Unmarshal(payload, &rawData)

	rawData.Metric = config.BloodPressure
	rawData.Unit = "mmHg"
	rawData.Timestamp = time.Now().Unix()
	fmt.Println(rawData)
}

// 血氧消息处理函数
var BloodOxygenMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()

	var rawData model.RawData
	json.Unmarshal(payload, &rawData)

	rawData.Metric = config.BloodOxygen
	rawData.Unit = "%"
	rawData.Timestamp = time.Now().Unix()
	fmt.Println(rawData)
}
