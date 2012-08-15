package main

import (
	"flag"
	"github.com/Krussell/asign"
	"github.com/Krussell/usb"
	"log"
)

var t = flag.String("t", "", "Template")
var a = flag.Bool("a", true, "Automatically allocate memory")

func main() {
	flag.Parse()
	if *t == "" {
		return
	}
	usb.Init()
	signUSB, err := openSignUSB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer signUSB.Close()

	sign := asign.New(signUSB)
	if *a == false {
		sign.DisableAutoMemory = true
	}

	sign.Reset()
	_, err = sign.WriteTemplate(*t)
	if err != nil {
		log.Println(err)
		return
	}
}
