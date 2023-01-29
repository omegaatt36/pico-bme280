package main

import (
	"fmt"
	"machine"

	"github.com/mailru/easyjson"
	"github.com/omegaatt36/pico-bme280/component/bme280/usecase"
	"github.com/omegaatt36/pico-bme280/present"
	"tinygo.org/x/drivers/bme280"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		SDA:       machine.GP0,
		SCL:       machine.GP1,
	})

	sensor := bme280.New(machine.I2C0)
	sensor.Configure()

	bme280UseCase := usecase.NewBME280UseCase()

	for {
		func() {
			var wait string
			n, err := fmt.Scanln(&wait)
			if n == 0 {
				return
			}

			if wait != "read" {
				fmt.Println(wait)
				return
			}

			o := bme280UseCase.Detect(&sensor)
			if o.Err != "" {
				fmt.Println(o.Err)
				return
			}

			jsonBME280 := present.JsonBME280{
				Err: o.Err,

				Temperature: o.Temperature,
				Pressure:    o.Pressure,
				Humidity:    o.Humidity,
			}

			bs, err := easyjson.Marshal(jsonBME280)
			if o.Err != "" {
				fmt.Println(err)
				return
			}

			fmt.Println(string(bs))
		}()
	}
}
