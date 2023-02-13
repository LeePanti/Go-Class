// Wait groups
package main

// Demonstrate flags

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//set flags
	msg := flag.String("msg", "Howdy stranger!", "the message to display")
	num := flag.Int("num", 1, "how many times to print the message")
	caps := flag.Bool("caps", false, "should the string be all caps")
	flag.Parse()

	for i := 0; i < *num; i++ {
		if *caps {
			fmt.Println(strings.ToUpper(*msg))
		} else {
			fmt.Println(*msg)
		}
	}
}
