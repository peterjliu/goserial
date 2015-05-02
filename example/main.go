// Program to read and log the Arduino serial output.
// Tested with arduinocode/temp36.ino loaded onto a Uno.
// This reports the temperature in celsius and voltage from a Temp36 sensor
// in CSV format.
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
	defer c.Close()
	check(err)
	for i := 0; i < 5; i++ {
		line, err := c.ReadLine()
		check(err)
		log.Print(line)
	}
}
