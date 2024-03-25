// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import "testing"

func TestGroupCharacters(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{`PW: A1c9.Ba`, `A1c9.Ba`, `ABca19.`},
		{`PW: PX4xDoiKrs,[egEAief{`, `PX4xDoiKrs,[egEAief{`, `PXDKEAxoirsegief4,[{`},
		{`PW: *Z%C9d+PZYkD7D+{~r'w`, `*Z%C9d+PZYkD7D+{~r'w`, `ZCPZYDDdkrw97*%++{~'`},
		{`PW: 4?r2YV:Abo&/z<3tJ*Z{`, `4?r2YV:Abo&/z<3tJ*Z{`, `YVAJZrbozt423?:&/<*{`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grouped := GroupCharsForMobile(tt.password)
			if grouped != tt.want {
				t.Errorf("GroupCharsForMobile() failed, expected: %s, got: %s", tt.want, grouped)
			}
		})
	}
}
