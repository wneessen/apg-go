// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

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
		{"ModeNumeric", ModeNumeric},
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
		{"ModeComplex", "C", []Mode{
			ModeLowerCase, ModeNumeric, ModeSpecial,
			ModeUpperCase,
		}},
		{"ModeHumanReadable", "H", []Mode{ModeHumanReadable}},
		{"ModeLowerCase", "L", []Mode{ModeLowerCase}},
		{"ModeNumeric", "N", []Mode{ModeNumeric}},
		{"ModeUpperCase", "U", []Mode{ModeUpperCase}},
		{"ModeSpecial", "S", []Mode{ModeSpecial}},
		{"ModeLowerSpecialUpper", "LSUH", []Mode{
			ModeHumanReadable,
			ModeLowerCase, ModeSpecial, ModeUpperCase,
		}},
		{"ModeComplexNoHumanReadable", "Ch", []Mode{
			ModeLowerCase,
			ModeNumeric, ModeSpecial, ModeUpperCase,
		}},
		{"ModeComplexNoLower", "Cl", []Mode{
			ModeNumeric, ModeSpecial,
			ModeUpperCase,
		}},
		{"ModeComplexNoNumber", "Cn", []Mode{
			ModeLowerCase, ModeSpecial,
			ModeUpperCase,
		}},
		{"ModeComplexNoSpecial", "Cs", []Mode{
			ModeLowerCase, ModeNumeric,
			ModeUpperCase,
		}},
		{"ModeComplexNoUpper", "Cu", []Mode{
			ModeLowerCase, ModeNumeric,
			ModeSpecial,
		}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mm := ModesFromFlags(tc.ms)
			for _, tm := range tc.mode {
				if !MaskHasMode(mm, tm) {
					t.Errorf("ModesFromFlags() failed, expected mode %q not found", tm)
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
		{"ModeNumeric", ModeNumeric, "Numeric"},
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
