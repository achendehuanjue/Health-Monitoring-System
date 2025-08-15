package model

// 原始数据结构
type RawData struct {
	DeviceID  string      // 设备ID
	Metric    string      // 指标类型
	Value     interface{} // 测量值
	Timestamp int64       // 时间戳(秒)
	Unit      string      // 单位
}
