package apg

import (
	"strings"
	"testing"
)

func TestConvertByteToWord(t *testing.T) {
	tests := []struct {
		name    string
		char    byte
		want    string
		wantErr bool
	}{
		{
			name:    "UpperCaseChar",
			char:    'A',
			want:    alphabetNames['A'],
			wantErr: false,
		},
		{
			name:    "LowerCaseChar",
			char:    'a',
			want:    strings.ToLower(alphabetNames['A']),
			wantErr: false,
		},
		{
			name:    "NonAlphaChar",
			char:    '(',
			want:    symbNumNames['('],
			wantErr: false,
		},
		{
			name:    "UnsupportedChar",
			char:    'ü',
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertByteToWord(tt.char)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertByteToWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertByteToWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
