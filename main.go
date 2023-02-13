// Wait groups
package main

// Demonstrate flags

import (
	"flag"
	"fmt"
)

func main() {
	//set flags
	msg := flag.String("msg", "Howdy stranger!", "the message to display")
	flag.Parse()

	// check if the user set the flag
	fmt.Println(*msg)
}
