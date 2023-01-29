package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
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
				return
			}
			defer port.Close()

			buff := make([]byte, 100)
			for {
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
					break
				}

				if n == 0 {
					continue
				}

				fmt.Printf("%v", string(buff[:n]))
			}
		}(portName)
	}

	<-ctx.Done()
}
