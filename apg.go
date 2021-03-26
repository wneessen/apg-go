package main

import (
	"fmt"
	"github.com/wneessen/apg.go/config"
	"log"
)

// Main function that generated the passwords and returns them
func main() {
	// Log config
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)

	// Create config
	conf := config.NewConfig()
	conf.ParseParams()
	pwLength := conf.GetPwLengthFromParams()
	charRange := getCharRange(conf)

	// Generate passwords
	for i := 1; i <= conf.NumOfPass; i++ {
		pwString, err := getRandChar(&charRange, pwLength)
		if err != nil {
			log.Fatalf("getRandChar returned an error: %q\n", err)
		}

		switch conf.OutputMode {
		case 1:
			{
				spelledPw, err := spellPasswordString(pwString)
				if err != nil {
					log.Fatalf("spellPasswordString returned an error: %q\n", err.Error())
				}
				fmt.Printf("%v (%v)\n", pwString, spelledPw)
				break
			}
		default:
			{
				fmt.Println(pwString)
				break
			}
		}
	}
}
