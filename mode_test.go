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
			var m ModeMask
			m = MaskSetMode(m, tc.mode)
			if !MaskHasMode(m, tc.mode) {
				t.Errorf("MaskSetMode() failed, mode not found in bitmask")
			}
			m = MaskToggleMode(m, tc.mode)
			if MaskHasMode(m, tc.mode) {
				t.Errorf("MaskToggleMode() failed, mode found in bitmask")
			}
			m = MaskToggleMode(m, tc.mode)
			if !MaskHasMode(m, tc.mode) {
				t.Errorf("MaskToggleMode() failed, mode not found in bitmask")
			}
			m = MaskClearMode(m, tc.mode)
			if MaskHasMode(m, tc.mode) {
				t.Errorf("MaskClearMode() failed, mode found in bitmask")
			}
		})
	}
}
