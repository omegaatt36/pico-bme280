package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/albenik/go-serial"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	mode := &serial.Mode{
		BaudRate: 115200,
	}

	portNames, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(portNames) == 0 {
		log.Fatal("No serial ports found!")
	}

	for _, portName := range portNames {
		fmt.Printf("Found port: %v\n", portName)
		go func(name string) {
			port, err := serial.Open(name, mode)
			if err != nil {
				log.Fatal(err)
			}

			for {
				var buff []byte
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
					break
				}

				if n == 0 {
					fmt.Println("\nEOF")
					continue
				}

				fmt.Printf("%v", string(buff[:n]))
			}
		}(portName)
	}

	<-ctx.Done()
}
