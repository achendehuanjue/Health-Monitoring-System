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

	// 使用带过滤的处理器
	handler := NewTemperatureHandler() // 或其他对应的处理器
	token := client.Subscribe("health/+/"+str, 1, handler)
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

// 使用过滤器的体温消息处理
func NewTemperatureHandler() mqtt.MessageHandler {
	// 示例：只接收36.0-42.0℃的体温数据
	filter := &RangeFilter{Min: 36.0, Max: 42.0}
	return FilterMessageHandler(filter, TemperatureMessageHandler)
}

// 使用过滤器的血糖消息处理
func NewBloodGlucoseHandler() mqtt.MessageHandler {
	// 示例：只接收70-180 mg/dL的血糖数据
	filter := &RangeFilter{Min: 70, Max: 180}
	return FilterMessageHandler(filter, BloodGlucoseMessageHandler)
}
