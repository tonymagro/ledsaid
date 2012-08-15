package main

import (
	"flag"
	"github.com/Krussell/asign"
	"github.com/Krussell/usb"
	"log"
)

var t = flag.String("t", "", "Template")
var a = flag.Bool("a", true, "Automatically allocate memory")
var c = flag.Bool("c", false, "Clear the sign")

func main() {
	flag.Parse()
	if *t == "" && *c == false {
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
	if *c {
		_, err = sign.WriteTemplate("{{Text .A}}{{.ETX}}")
		if err != nil {
			log.Println(err)
		}
		return
	}
	_, err = sign.WriteTemplate(*t)
	if err != nil {
		log.Println(err)
		return
	}
}
