package apg

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	if c == nil {
		t.Errorf("NewConfig() failed, expected config pointer but got nil")
		return
	}
	c = NewConfig(nil)
	if c == nil {
		t.Errorf("NewConfig() failed, expected config pointer but got nil")
		return
	}
	if c.MinLength != DefaultMinLength {
		t.Errorf("NewConfig() failed, expected min length: %d, got: %d", DefaultMinLength,
			c.MinLength)
	}
	if c.MaxLength != DefaultMaxLength {
		t.Errorf("NewConfig() failed, expected max length: %d, got: %d", DefaultMaxLength,
			c.MaxLength)
	}
	if c.NumberPass != DefaultNumberPass {
		t.Errorf("NewConfig() failed, expected number of passwords: %d, got: %d",
			DefaultNumberPass, c.NumberPass)
	}
}

func TestWithMaxLength(t *testing.T) {
	var e int64 = 123
	c := NewConfig(WithMaxLength(e))
	if c == nil {
		t.Errorf("NewConfig(WithMaxLength()) failed, expected config pointer but got nil")
		return
	}
	if c.MaxLength != e {
		t.Errorf("NewConfig(WithMaxLength()) failed, expected max length: %d, got: %d",
			e, c.MaxLength)
	}
}

func TestWithMinLength(t *testing.T) {
	var e int64 = 1
	c := NewConfig(WithMinLength(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinLength()) failed, expected config pointer but got nil")
		return
	}
	if c.MinLength != e {
		t.Errorf("NewConfig(WithMinLength()) failed, expected min length: %d, got: %d",
			e, c.MinLength)
	}
}

func TestWithNumberPass(t *testing.T) {
	var e int64 = 123
	c := NewConfig(WithNumberPass(e))
	if c == nil {
		t.Errorf("NewConfig(WithNumberPass()) failed, expected config pointer but got nil")
		return
	}
	if c.NumberPass != e {
		t.Errorf("NewConfig(WithNumberPass()) failed, expected number of passwords: %d, got: %d",
			e, c.NumberPass)
	}
}
