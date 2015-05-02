// Program to read and log the Arduino serial output.
package main

import (
	"log"

	"github.com/peterjliu/goserial"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c, err := goserial.NewClient("/dev/cu.usbmodem1421", 9600)
	check(err)
	for {
		line, err := c.ReadLine()
		check(err)
		log.Print(line)
	}
}
