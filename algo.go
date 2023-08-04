package apg

// Algorithm is a type wrapper for an int type to represent different
// password generation algorithm
type Algorithm int

const (
	// Pronouncable represents the algorithm for pronouncable passwords
	// (koremutake syllables)
	Pronouncable Algorithm = iota
	// Random represents the algorithm for purely random passwords according
	// to the provided password modes/flags
	Random
	// Unsupported represents an unsupported algorithm
	Unsupported
)

// IntToAlgo takes an int value as input and returns the corresponding
// Algorithm
func IntToAlgo(a int) Algorithm {
	switch a {
	case 0:
		return Pronouncable
	case 1:
		return Random
	default:
		return Unsupported
	}
}
