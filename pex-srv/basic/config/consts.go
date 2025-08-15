package config

const (
	Temperature   = "temperature"    // 体温
	BloodGlucose  = "blood_glucose"  // 血糖
	BloodPressure = "blood_pressure" // 血压
	BloodOxygen   = "blood_oxygen"   // 血氧
)

// 各类数据的判断标准范围常量定义
const (
	// 体温标准范围(℃)
	TemperatureMin   = 35.0
	TemperatureMax   = 42.0
	TemperatureDelta = 1.0 // 短时间波动阈值

	// 血糖标准范围(mmol/L)
	BloodGlucoseMin   = 2.0
	BloodGlucoseMax   = 20.0
	FastingGlucoseMin = 3.9
	FastingGlucoseMax = 6.1
	PostprandialMax   = 11.1 // 餐后血糖上限

	// 血压标准范围(mmHg)
	SystolicMin      = 50
	SystolicMax      = 220
	DiastolicMin     = 40
	DiastolicMax     = 130
	PulsePressureMin = 20 // 脉压差下限
	PulsePressureMax = 60 // 脉压差上限
	SystolicDelta    = 30 // 收缩压波动阈值
	DiastolicDelta   = 20 // 舒张压波动阈值

	// 血氧标准范围(%)
	BloodOxygenMin   = 80.0
	BloodOxygenMax   = 100.0
	BloodOxygenLow   = 90.0 // 低血氧阈值
	BloodOxygenDelta = 5.0  // 短时间波动阈值

	// 时间相关阈值(秒)
	TimestampValidRange     = 1800 // 时间戳有效范围(±30分钟)
	TemperatureTimeWindow   = 1800 // 体温波动时间窗口(30分钟)
	BloodPressureTimeWindow = 3600 // 血压波动时间窗口(1小时)
	BloodOxygenTimeWindow   = 10   // 血氧波动时间窗口(10秒)
)
