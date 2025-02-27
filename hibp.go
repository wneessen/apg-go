// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import (
	"fmt"
	"time"

	"github.com/wneessen/go-hibp"
)

// HasBeenPwned checks the given password string against the HIBP pwned
// passwords database and returns true if the password has been leaked
func HasBeenPwned(password string) (bool, error) {
	hc := hibp.New(hibp.WithHTTPTimeout(time.Second*2),
		hibp.WithPwnedPadding())
	matches, _, err := hc.PwnedPassAPI.CheckPassword(password)
	if err != nil {
		return false, fmt.Errorf("failed to check password against HIBP: %w", err)
	}
	return matches.Present() && matches.Count > 0, err
}
