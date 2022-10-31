package hw02unpackstring

import (
	"errors"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "d0abc", expected: "abc"},
		{input: "d2abc", expected: "ddabc"},
		// uncomment if task with asterisk completed
		{input: `\4\5`, expected: `45`},
		{input: `\\4`, expected: `\\\\`},
		{input: `\\a`, expected: `\a`},
		{input: `\\\\4`, expected: `\\\\\`},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `qwe\\5a`, expected: `qwe\\\\\a`},
		{input: `🦍5`, expected: `🦍🦍🦍🦍🦍`},
		{input: `🦡\52\\\4b0b1🦡`, expected: `🦡55\4b🦡`},
		{input: `\\🦍-gorilla, 🦡-badger\\`, expected: `\🦍-gorilla, 🦡-badger\`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `qw\ne`, `\\\a`, `\455`, `\🦍`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
