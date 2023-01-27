package present

import "github.com/omegaatt36/pico-bme280/domain"

type JsonBME280 struct {
	Err string `json:"err,omitempty"`

	Temperature int32 `json:"temperature"`
	Pressure    int32 `json:"pressure"`
	Humidity    int32 `json:"humidity"`
}

func (bme280 *JsonBME280) ToBME280() *domain.BME280 {
	return nil
}
