package main

import (
	"fmt"
	"github.com/tonymagro/asign"
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
		fmt.Printf("{Hold}{%s}%s\n", k, k)
	}
	fmt.Println("{ETX}")

	fmt.Println("{STX}\n{WriteText}{B}")
	for k := range asign.ExtendedCharacter {
		fmt.Printf("{RollUp}{DimRed}%s:{AutoColor}{%s}\n", k, k)
	}
	fmt.Println("{ETX}")

	fmt.Println("{STX}\n{WriteText}{C}")
	for k := range asign.SpecialMode {
		fmt.Printf("{%s}%s\n", k, k)
	}
	fmt.Println("{ETX}")

	fmt.Println("{STX}\n{WriteText}{D}")
	for k := range asign.SpecialGraphics {
		fmt.Printf("{Hold}{DimGreen}%s{AutoColor}{%s}\n", k, k)
	}
	fmt.Println("{ETX}")

	fmt.Println("{EOT}")
}
