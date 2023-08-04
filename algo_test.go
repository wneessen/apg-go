package apg

import "testing"

func TestIntToAlgo(t *testing.T) {
	tt := []struct {
		name string
		a    int
		e    Algorithm
	}{
		{"AlgoPronouncable", 0, AlgoPronouncable},
		{"AlgoRandom", 1, AlgoRandom},
		{"AlgoUnsupported", 2, AlgoUnsupported},
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
