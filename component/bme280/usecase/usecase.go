package usecase

import "github.com/omegaatt36/pico-bme280/domain"

type bme280UseCase struct {
}

func NewBME280UseCase() domain.BME280UseCase {
	return &bme280UseCase{}
}

func (o *bme280UseCase) Detect(sensor domain.SensorBME280) *domain.BME280 {
	bme280 := domain.BME280{}

	temperature, err := sensor.ReadTemperature()
	if err != nil {
		bme280.Err = err.Error()
		return &bme280
	}
	bme280.Temperature = temperature

	pressure, err := sensor.ReadPressure()
	if err != nil {
		bme280.Err = err.Error()
		return &bme280
	}
	bme280.Pressure = pressure

	humidity, err := sensor.ReadHumidity()
	if err != nil {
		bme280.Err = err.Error()
		return &bme280
	}
	bme280.Humidity = humidity

	return &bme280
}
