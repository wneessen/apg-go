// Package main is the APG command line client that makes use of the apg-go library

package main

import (
	"fmt"
	"os"

	"github.com/wneessen/apg-go"
)

func main() {
	g := apg.New()
	rb, err := g.RandomBytes(8)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	fmt.Printf("Random: %#v\n", rb)
}
