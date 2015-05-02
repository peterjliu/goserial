// Client for reading serial device output (e.g. arduino)
package goserial

import (
	"bufio"
	"fmt"

	"github.com/tarm/serial"
)

type Client struct {
	port *serial.Port
	open bool
}

// Opens device ready to be read
func NewClient(dev string, baud int) (*Client, error) {
	c := &serial.Config{Name: "/dev/cu.usbmodem1421", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	cl := &Client{
		port: s,
		open: true,
	}
	return cl, nil
}

// Reads next line after the first newline encountered
func (c *Client) ReadLine() (string, error) {
	if !c.open {
		return "", fmt.Errorf("Serial port closed.\n")
	}
	reader := bufio.NewReader(c.port)
	// discard up to first newline
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line, nil
}

func (c *Client) Close() {
	c.port.Close()
}
