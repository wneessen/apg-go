// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import "testing"

func TestIntToAlgo(t *testing.T) {
	tt := []struct {
		name string
		a    int
		e    Algorithm
	}{
		{"AlgoPronounceable", 0, AlgoPronounceable},
		{"AlgoRandom", 1, AlgoRandom},
		{"AlgoCoinflip", 2, AlgoCoinFlip},
		{"AlgoUnsupported", 3, AlgoUnsupported},
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
