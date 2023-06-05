package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/albenik/go-serial/v2"
	"github.com/omegaatt36/pico-bme280/present"
)

type serialBME280 struct {
	temperature int32
	humidity    int32
	pressure    int32

	port *serial.Port
}

func newSerialBME280() (*serialBME280, error) {
	ops := []serial.Option{
		serial.WithBaudrate(115200),
	}

	portNames, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	if len(portNames) == 0 {
		return nil, errors.New("No serial ports found!")
	}

	var bme280 serialBME280

	for _, portName := range portNames {
		fmt.Printf("Found port: %v\n", portName)
		port, err := serial.Open(portName, ops...)
		if err != nil {
			return nil, err
		}

		bme280.port = port
		break
	}

	return &bme280, nil
}

func (bme280 *serialBME280) read() error {
	if _, err := bme280.port.Write([]byte("read\n\r")); err != nil {
		return err
	}

	buff := make([]byte, 59)
	n, err := bme280.port.Read(buff)
	if err != nil {
		return err
	}

	if n == 0 {
		return nil
	}

	read := bytes.Trim(buff[:n], "\n")

	var jsonBME280 present.JsonBME280
	if err := json.Unmarshal(read, &jsonBME280); err != nil {
		return err
	}

	bme280.temperature = jsonBME280.Temperature
	bme280.humidity = jsonBME280.Humidity
	bme280.pressure = jsonBME280.Pressure

	return nil
}
