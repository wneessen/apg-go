package apg

import (
	"testing"
)

func TestSetClearHasToggleMode(t *testing.T) {
	tt := []struct {
		name string
		mode Mode
	}{
		{"ModeNumber", ModeNumber},
		{"ModeLowerCase", ModeLowerCase},
		{"ModeUpperCase", ModeUpperCase},
		{"ModeSpecial", ModeSpecial},
		{"ModeHumanReadable", ModeHumanReadable},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var m Mode
			m = SetMode(m, tc.mode)
			if !HasMode(m, tc.mode) {
				t.Errorf("SetMode() failed, mode not found in bitmask")
			}
			m = ToggleMode(m, tc.mode)
			if HasMode(m, tc.mode) {
				t.Errorf("ToggleMode() failed, mode found in bitmask")
			}
			m = ToggleMode(m, tc.mode)
			if !HasMode(m, tc.mode) {
				t.Errorf("ToggleMode() failed, mode not found in bitmask")
			}
			m = ClearMode(m, tc.mode)
			if HasMode(m, tc.mode) {
				t.Errorf("ClearMode() failed, mode found in bitmask")
			}
		})
	}
}
