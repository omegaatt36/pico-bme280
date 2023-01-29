package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/albenik/go-serial/v2"
	"github.com/omegaatt36/pico-bme280/present"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type serialBME280 struct {
	temperature int32
	humidity    int32
	pressure    int32

	port *serial.Port
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

var (
	addr = flag.String("listen-address", ":9110", "The address to listen on for HTTP requests.")
)

type bme280Collector struct {
	temperatureMetric *prometheus.Desc
	humidityMetric    *prometheus.Desc
	pressureMetric    *prometheus.Desc

	sensor *serialBME280
}

func newBme280Collector(bme280 *serialBME280) *bme280Collector {
	return &bme280Collector{
		temperatureMetric: prometheus.NewDesc(
			"bme280_temperature", "bme 280 temperature(Celsius)",
			nil, nil),
		humidityMetric: prometheus.NewDesc(
			"bme280_humidity", "bme 280 humidity(RH)",
			nil, nil),
		pressureMetric: prometheus.NewDesc(
			"bme280_pressure", "bme 280 pressure(hPa)",
			nil, nil),

		sensor: bme280,
	}
}

// Each and every collector must implement the Describe function.
// It essentially writes all descriptors to the prometheus desc channel.
func (collector *bme280Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.temperatureMetric
	ch <- collector.pressureMetric
	ch <- collector.humidityMetric
}

// Collect implements required collect function for all promehteus collectors
func (collector *bme280Collector) Collect(ch chan<- prometheus.Metric) {
	if err := collector.sensor.read(); err != nil {
		log.Println(err)
	}

	mTemperature := prometheus.MustNewConstMetric(collector.temperatureMetric,
		prometheus.GaugeValue, float64(collector.sensor.temperature)/100)
	mHumidity := prometheus.MustNewConstMetric(collector.humidityMetric,
		prometheus.GaugeValue, float64(collector.sensor.humidity)/100)
	mPressure := prometheus.MustNewConstMetric(collector.pressureMetric,
		prometheus.GaugeValue, float64(collector.sensor.pressure)/100000)
	ch <- prometheus.NewMetricWithTimestamp(time.Now(), mTemperature)
	ch <- prometheus.NewMetricWithTimestamp(time.Now(), mHumidity)
	ch <- prometheus.NewMetricWithTimestamp(time.Now(), mPressure)
}

func main() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ops := []serial.Option{
		serial.WithBaudrate(115200),
	}

	portNames, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(portNames) == 0 {
		log.Fatal("No serial ports found!")
	}

	var bme280 serialBME280

	for _, portName := range portNames {
		fmt.Printf("Found port: %v\n", portName)
		port, err := serial.Open(portName, ops...)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer port.Close()

		bme280.port = port
		if err := bme280.read(); err != nil {
			log.Fatal(err)
		}
	}
	collector := newBme280Collector(&bme280)
	prometheus.MustRegister(collector)

	log.Printf("handle %s/metrics \n", *addr)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

	log.Println("exit")

	<-ctx.Done()
}
