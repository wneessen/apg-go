// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

// Algorithm is a type wrapper for an int type to represent different
// password generation algorithm
type Algorithm int

const (
	// AlgoPronounceable represents the algorithm for pronounceable passwords
	// (koremutake syllables)
	AlgoPronounceable Algorithm = iota
	// AlgoRandom represents the algorithm for purely random passwords according
	// to the provided password modes/flags
	AlgoRandom
	// AlgoCoinFlip represents a very simple coinflip algorithm returning "heads"
	// or "tails"
	AlgoCoinFlip
	// AlgoBinary represents a full binary randomness mode with up to 256 bits
	// of randomness
	AlgoBinary
	// AlgoUnsupported represents an unsupported algorithm
	AlgoUnsupported
)

// IntToAlgo takes an int value as input and returns the corresponding
// Algorithm
func IntToAlgo(a int) Algorithm {
	switch a {
	case 0:
		return AlgoPronounceable
	case 1:
		return AlgoRandom
	case 2:
		return AlgoCoinFlip
	case 3:
		return AlgoBinary
	default:
		return AlgoUnsupported
	}
}
