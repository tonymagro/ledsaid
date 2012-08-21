package main

import (
	"encoding/json"
	"flag"
	"github.com/krussell/asign"
	"github.com/krussell/usb"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	t     = flag.String("t", "", "Template")
	a     = flag.Bool("a", true, "Automatically allocate memory")
	b     = flag.Bool("b", false, "Blank the sign")
	s     = flag.Bool("s", false, "Read from stdin")
	f     = flag.String("f", "", "Read from file")
	w     = flag.String("w", "", "Read from web address")
	bjson = flag.Bool("json", false, "Export protocol to json")
)

func main() {
	flag.Parse()

	if *bjson {
		j, err := json.MarshalIndent(&protocol, "", "   ")
		if err != nil {
			log.Fatalln(err)
			return
		}
		println(string(j))
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

	if *w != "" {
		println("Reading from web address:", *w)
		r, err := http.Get(*w)
		if err != nil {
			log.Println(err)
			return
		}
		buf, err := ioutil.ReadAll(r.Body)
		println(string(buf))
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
