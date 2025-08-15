package pkg

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"pex-srv/handler/model"
	"time"
)

// 数据过滤器接口
type DataFilter interface {
	Filter(rawData *model.RawData) bool
}

// 范围过滤器
type RangeFilter struct {
	Min float64
	Max float64
}

func (f *RangeFilter) Filter(rawData *model.RawData) bool {
	value, ok := rawData.Value.(float64)
	if !ok {
		return false
	}
	return value >= f.Min && value <= f.Max
}

// 时间过滤器
type TimeFilter struct {
	StartTime time.Time
	EndTime   time.Time
}

func (f *TimeFilter) Filter(rawData *model.RawData) bool {
	t := time.Unix(rawData.Timestamp, 0)
	return t.After(f.StartTime) && t.Before(f.EndTime)
}

// 复合过滤器
type CompositeFilter struct {
	Filters []DataFilter
}

func (f *CompositeFilter) Filter(rawData *model.RawData) bool {
	for _, filter := range f.Filters {
		if !filter.Filter(rawData) {
			return false
		}
	}
	return true
}

// 过滤消息处理器
func FilterMessageHandler(filter DataFilter, handler mqtt.MessageHandler) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		var rawData model.RawData
		json.Unmarshal(payload, &rawData)

		if filter.Filter(&rawData) {
			handler(client, msg)
		}
	}
}
