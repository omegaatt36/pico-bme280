package domain

// BME280Reader defines bme280 reader
type SensorBME280 interface {
	ReadTemperature() (int32, error)
	ReadPressure() (int32, error)
	ReadHumidity() (int32, error)
}

// BME280 defines probe bme280.
type BME280 struct {
	Err string

	Temperature int32
	Pressure    int32
	Humidity    int32
}

type BME280UseCase interface {
	Detect(SensorBME280) *BME280
}
