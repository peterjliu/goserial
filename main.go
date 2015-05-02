// Program to read and log the Arduino serial output.
package main

import (
	"bufio"
	"log"

	"github.com/tarm/serial"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := &serial.Config{Name: "/dev/cu.usbmodem1421", Baud: 9600}
	s, err := serial.OpenPort(c)
	check(err)

	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadString('\n')
		check(err)
		log.Print(line)
	}
}
