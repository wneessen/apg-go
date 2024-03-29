// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

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
				t.Errorf("ConvertByteToWord() error = %s, wantErr %t", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertByteToWord() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSpell(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "a",
			want:    "alfa",
			wantErr: false,
		},
		{
			name:    "multiple characters",
			input:   "abc",
			want:    "alfa/bravo/charlie",
			wantErr: false,
		},
		{
			name:    "non-alphabetic character",
			input:   "1",
			want:    "ONE",
			wantErr: false,
		},
		{
			name:    "mixed alphabetic and non-alphabetic characters",
			input:   "a1",
			want:    "alfa/ONE",
			wantErr: false,
		},
		{
			name:    "not supported characters",
			input:   "üäöß€",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Spell(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spell() error = %s, wantErr %t", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Spell() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestPronounce(t *testing.T) {
	tests := []struct {
		name      string
		syllables []string
		want      string
		wantErr   bool
	}{
		{
			name:      "Pronounce_NoSyllables",
			syllables: []string{},
			want:      "",
			wantErr:   false,
		},
		{
			name:      "Pronounce_SingleSyllable",
			syllables: []string{"me"},
			want:      "me",
			wantErr:   false,
		},
		{
			name:      "Pronounce_MultipleSyllables",
			syllables: []string{"mu", "sa"},
			want:      "mu-sa",
			wantErr:   false,
		},
		{
			name:      "Pronounce_Mixed",
			syllables: []string{"mu", "1"},
			want:      "mu-ONE",
			wantErr:   false,
		},
		{
			name:      "Pronounce_NonKoremutakeSyllable",
			syllables: []string{"ä"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()
			g := New(config)
			g.syllables = tt.syllables
			got, err := g.Pronounce()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.Pronounce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generator.Pronounce() = %v, want %v", got, tt.want)
			}
		})
	}
}
