package main

import (
	"fmt"
	"github.com/wneessen/apg-go/chars"
	"github.com/wneessen/apg-go/config"
	"github.com/wneessen/apg-go/random"
)

func main() {
	c := config.Config{
		UseNumber:    true,
		UseSpecial:   true,
		UseUpperCase: true,
		UseLowerCase: true,
		PwAlgo:       1,
		MinPassLen:   15,
		MaxPassLen:   15,
	}
	pl := config.GetPwLengthFromParams(&c)
	cs := chars.GetRange(&c)
	pw, err := random.GetChar(cs, pl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your Password:", pw)
}
