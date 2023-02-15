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
	rev := flag.Bool("r", false, "reverses the string")
	flag.Parse()

	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	if *rev {
		var message string
		for _, i := range *msg {
			message = string(i) + message
		}
		*msg = message
	}

	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}

}
