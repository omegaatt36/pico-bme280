package main

import (
	"context"
	"log"
	"net/http"

	"github.com/omegaatt36/pico-bme280/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

var listenAddress string

func Main(ctx context.Context) {
	bme280, err := newSerialBME280()
	if err != nil {
		log.Fatal(err)
	}

	collector := newBme280Collector(bme280)
	prometheus.MustRegister(collector)

	server := &http.Server{Addr: listenAddress}

	go func() {
		log.Printf("handle %s/metrics \n", listenAddress)
		http.Handle("/metrics", promhttp.Handler())
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP Server Shutdown Error: %v", err)
	}
	log.Println("exit")
}

func main() {
	app := app.App{
		Main:  Main,
		Flags: []cli.Flag{},
	}

	app.Flags = append(app.Flags,
		&cli.StringFlag{
			Name:        "listen-address",
			EnvVars:     []string{"LISTEN_ADDRESS"},
			DefaultText: ":9110",
			Value:       ":9110",
			Usage:       "The address to listen on for HTTP requests.",
			Destination: &listenAddress,
		},
	)

	app.Run()
}
