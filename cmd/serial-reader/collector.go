package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
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
		return
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
