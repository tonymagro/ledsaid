package main

import (
	"errors"
	"fmt"
	"github.com/tonymagro/asign"
	"github.com/tonymagro/usb"
	"log"
)

const (
	ALPHA_VENDOR_ID    = 0x8765
	ALPHA_PRODUCT_ID   = 0x1234
	ALPHA_INTERFACE_ID = 0x00
)
const (
	RECV_LENGTH    = 0x40 // Max Packet Size of 64 
	SEND_LENGTH    = 0x40 // Max Packet Size of 64 
	WRITE_ENDPOINT = 0x02
	READ_ENDPOINT  = 0x82
	WRITE_TIMEOUT  = 5000
	READ_TIMEOUT   = 5000
)

type SignUSB usb.Device

func (self *SignUSB) Write(p []byte) (n int, err error) {
	fmt.Println(asign.PacketString(p))
	dev := usb.Device(*self)
	dev.BulkWrite(WRITE_ENDPOINT, p)
	n = len(p)
	return
}

func (self *SignUSB) Read(p []byte) (n int, err error) {
	return
}

func (self *SignUSB) Close() {
	dev := usb.Device(*self)
	v,p := dev.Vendor(), dev.Product()
	dev.Close()
	log.Println("Closed Device: ", v, p)
}

func openSignUSB() (sign *SignUSB, err error) {
	dev := usb.Open(ALPHA_VENDOR_ID, ALPHA_PRODUCT_ID)
	if dev == nil {
		err = errors.New("Invalid Device")
		return
	}

	dev.Interface(0)
	//dev.timeout = WRITE_TIMEOUT
	log.Println("Opened Device: ", dev.Vendor(), dev.Product())
	s := SignUSB(*dev)
	return &s, nil
}
