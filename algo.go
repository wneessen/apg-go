package apg

// Algorithm is a type wrapper for an int type to represent different
// password generation algorithm
type Algorithm int

const (
	// AlgoPronouncable represents the algorithm for pronouncable passwords
	// (koremutake syllables)
	AlgoPronouncable Algorithm = iota
	// AlgoRandom represents the algorithm for purely random passwords according
	// to the provided password modes/flags
	AlgoRandom
	// AlgoCoinFlip represents a very simple coinflip algorithm returning "heads"
	// or "tails"
	AlgoCoinFlip
	// AlgoUnsupported represents an unsupported algorithm
	AlgoUnsupported
)

// IntToAlgo takes an int value as input and returns the corresponding
// Algorithm
func IntToAlgo(a int) Algorithm {
	switch a {
	case 0:
		return AlgoPronouncable
	case 1:
		return AlgoRandom
	case 2:
		return AlgoCoinFlip
	default:
		return AlgoUnsupported
	}
}
