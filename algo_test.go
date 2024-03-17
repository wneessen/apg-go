// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import (
	"testing"
)

func TestIntToAlgo(t *testing.T) {
	tt := []struct {
		name string
		a    int
		e    Algorithm
	}{
		{"AlgoPronounceable", 0, AlgoPronounceable},
		{"AlgoRandom", 1, AlgoRandom},
		{"AlgoCoinflip", 2, AlgoCoinFlip},
		{"AlgoBinary", 3, AlgoBinary},
		{"AlgoUnsupported", 4, AlgoUnsupported},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a := IntToAlgo(tc.a)
			if a != tc.e {
				t.Errorf("IntToAlgo() failed, expected: %d, got: %d", tc.e, a)
			}
		})
	}
}

func FuzzIntToAlgo(f *testing.F) {
	f.Add(-1)  // Test negative input
	f.Add(4)   // Test out-of-range positive input
	f.Add(100) // Test very large input
	f.Fuzz(func(t *testing.T, a int) {
		algo := IntToAlgo(a)
		switch a {
		case 0:
			if algo != AlgoPronounceable {
				t.Errorf("IntToAlgo(%d) expected AlgoPronounceable, got %v", a, algo)
			}
		case 1:
			if algo != AlgoRandom {
				t.Errorf("IntToAlgo(%d) expected AlgoRandom, got %v", a, algo)
			}
		case 2:
			if algo != AlgoCoinFlip {
				t.Errorf("IntToAlgo(%d) expected AlgoCoinFlip, got %v", a, algo)
			}
		case 3:
			if algo != AlgoBinary {
				t.Errorf("IntToAlgo(%d) expected AlgoBinary, got %v", a, algo)
			}
		default:
			if algo != AlgoUnsupported {
				t.Errorf("IntToAlgo(%d) expected AlgoUnsupported, got %v", a, algo)
			}
		}
	})
}
