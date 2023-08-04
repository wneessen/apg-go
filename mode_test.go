package apg

import (
	"testing"
)

func TestSetClearHasToggleMode(t *testing.T) {
	tt := []struct {
		name string
		mode Mode
	}{
		{"ModeHumanReadable", ModeHumanReadable},
		{"ModeLowerCase", ModeLowerCase},
		{"ModeNumber", ModeNumber},
		{"ModeSpecial", ModeSpecial},
		{"ModeUpperCase", ModeUpperCase},
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

func TestModesFromFlags(t *testing.T) {
	tt := []struct {
		name string
		ms   string
		mode []Mode
	}{
		{"ModeComplex", "C", []Mode{ModeLowerCase, ModeNumber, ModeSpecial,
			ModeUpperCase}},
		{"ModeHumanReadable", "H", []Mode{ModeHumanReadable}},
		{"ModeLowerCase", "L", []Mode{ModeLowerCase}},
		{"ModeNumber", "N", []Mode{ModeNumber}},
		{"ModeUpperCase", "U", []Mode{ModeUpperCase}},
		{"ModeSpecial", "S", []Mode{ModeSpecial}},
		{"ModeLowerSpecialUpper", "LSUH", []Mode{ModeHumanReadable,
			ModeLowerCase, ModeSpecial, ModeUpperCase}},
		{"ModeComplexNoHumanReadable", "Ch", []Mode{ModeLowerCase,
			ModeNumber, ModeSpecial, ModeUpperCase}},
		{"ModeComplexNoLower", "Cl", []Mode{ModeNumber, ModeSpecial,
			ModeUpperCase}},
		{"ModeComplexNoNumber", "Cn", []Mode{ModeLowerCase, ModeSpecial,
			ModeUpperCase}},
		{"ModeComplexNoSpecial", "Cs", []Mode{ModeLowerCase, ModeNumber,
			ModeUpperCase}},
		{"ModeComplexNoUpper", "Cu", []Mode{ModeLowerCase, ModeNumber,
			ModeSpecial}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var mm ModeMask
			mm = ModesFromFlags(tc.ms)
			for _, tm := range tc.mode {
				if !MaskHasMode(mm, tm) {
					t.Errorf("ModesFromFlags() failed, expected mode %q not found",
						tm)
				}
			}
		})
	}
}

func TestMode_String(t *testing.T) {
	tt := []struct {
		name string
		m    Mode
		e    string
	}{
		{"ModeHumanReadable", ModeHumanReadable, "Human-readable"},
		{"ModeLowerCase", ModeLowerCase, "Lower-case"},
		{"ModeNumber", ModeNumber, "Number"},
		{"ModeSpecial", ModeSpecial, "Special"},
		{"ModeUpperCase", ModeUpperCase, "Upper-case"},
		{"ModeUnknown", 255, "Unknown"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.m.String() != tc.e {
				t.Errorf("Mode.String() failed, expected: %s, got: %s", tc.e,
					tc.m.String())
			}
		})
	}
}
