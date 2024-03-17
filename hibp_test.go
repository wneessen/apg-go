// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import (
	"testing"
)

func TestHasBeenPwned(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"Pwned PW", "Test123", true},
		{"Secure PW", "Cta8mWYmW7O*j1V!YMTS", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HasBeenPwned(tt.password)
			if err != nil {
				t.Logf("HasBeenPwned() failed: %s", err)
				return
			}
			if tt.want != got {
				t.Errorf("HasBeenPwned() failed, wanted: %t, got: %t", tt.want, got)
			}
		})
	}
}
