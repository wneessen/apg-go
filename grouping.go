// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import "unicode"

// GroupCharsForMobile takes a given string of characters and groups them in a mobile-friendly
// manner. The grouping is based on the following precedense: uppercase, lowercase, numbers
// and special characters. The grouped string is then returned.
func GroupCharsForMobile(chars string) string {
	var uppers, lowers, numbers, others []rune
	for _, char := range chars {
		switch {
		case unicode.IsUpper(char):
			uppers = append(uppers, char)
		case unicode.IsLower(char):
			lowers = append(lowers, char)
		case unicode.IsNumber(char):
			numbers = append(numbers, char)
		default:
			others = append(others, char)
		}
	}
	return string(uppers) + string(lowers) + string(numbers) + string(others)
}
