package main

import (
	"fmt"
	"github.com/Krussell/asign"
)

func main() {
	fmt.Println("{SOT}")
	fmt.Println("{STX}\n{WriteText}{A}")

	for k := range asign.ModeCode {
		if k == "Special" {
			continue
		}
		fmt.Printf("{%s}%s\n", k, k)
	}

	for k := range asign.Color {
		fmt.Printf("{WipeIn}{%s}%s\n", k, k)
	}

	fmt.Println("{ETX}")
	fmt.Println("{EOT}")
}
