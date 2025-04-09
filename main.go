package main

import (
	"flag"
	"fmt"
)

func main() {
	var chFlag = flag.Int("ch", 0, "which chapter to run")

	flag.Parse()

	switch *chFlag {
	case 5:
		Ch5()
	case 6:
		Ch6()
	case 7:
		Ch7()
	default:
		fmt.Println("Unknown chapter: ", *chFlag)
	}
}
