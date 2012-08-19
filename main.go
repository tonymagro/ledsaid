package main

import (
	"flag"
	"github.com/Krussell/asign"
	"github.com/Krussell/usb"
	"io/ioutil"
	"log"
	"os"
)

var (
	t = flag.String("t", "", "Template")
	a = flag.Bool("a", true, "Automatically allocate memory")
	b = flag.Bool("b", false, "Blank the sign")
	s = flag.Bool("s", false, "Read from stdin")
	f = flag.String("f", "", "Read from file")
)

func main() {
	flag.Parse()
	if *t == "" && *b == false && *s == false && *f == "" {
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
	if *b {
		sign.DisableAutoMemory = false
		_, err = sign.WriteTemplate([]byte("{SOT}{STX}{WriteText}{A}{ETX}{EOT}"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if *t != "" {
		_, err = sign.WriteTemplate([]byte(*t))
		if err != nil {
			log.Println(err)
			return
		}
	}

	if *f != "" {
		f, err := os.Open(*f)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		buf, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sign.WriteTemplate(buf)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if *s {
		println("Reading from stdin")
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sign.WriteTemplate(buf)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
