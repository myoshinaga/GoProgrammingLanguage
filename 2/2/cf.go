package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/myoshinaga/GoProgrammingLanguage/2/2/conv"
)

var c = flag.String("c", "tc", "convert")

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		str := flag.Arg(0)
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		printConvertValue(*c, f)
	} else {

	}
}

func printConvertValue(opt string, f float64) {
	switch opt {
	case "tc":
		r := conv.Celsius(f)
		fmt.Printf("Celsius to Fahrenheit: %s = %s\n", r, conv.CtoF(r))
	case "tf":
		r := conv.Fahrenheit(f)
		fmt.Printf("Fahrenheit to Celsius: %s = %s\n", r, conv.FtoC(r))
	case "lf":
		r := conv.Feet(f)
		fmt.Printf("Feet to Meters: %s = %s\n", r, conv.FtoM(r))
	case "lm":
		r := conv.Meters(f)
		fmt.Printf("Meters to Feet: %s = %s\n", r, conv.MtoF(r))
	case "wp":
		r := conv.Pondus(f)
		fmt.Printf("Pondus to Kilogram: %s = %s\n", r, conv.PtoK(r))
	case "wk":
		r := conv.Kilogram(f)
		fmt.Printf("Kilogram to Pondus: %s = %s\n", r, conv.KtoP(r))
	}
}
