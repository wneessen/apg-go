package apg

// Generator is the password generator type of the APG package
type Generator struct {
	// charRange is the range of character used for the
	charRange string
}

// New returns a new password Generator type
func New() *Generator {
	return &Generator{}
}
