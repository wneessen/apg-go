package apg

import (
	"time"

	"github.com/wneessen/go-hibp"
)

// HasBeenPwned checks the given password string against the HIBP pwned
// passwords database and returns true if the password has been leaked
func HasBeenPwned(password string) (bool, error) {
	hc := hibp.New(hibp.WithHTTPTimeout(time.Second*2),
		hibp.WithPwnedPadding())
	matches, _, err := hc.PwnedPassAPI.CheckPassword(password)
	return matches != nil && matches.Count != 0, err
}
