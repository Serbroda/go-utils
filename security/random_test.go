package security

import (
	"testing"
	"unicode"
)

func TestGenerateRandomAlphabeticString_Length(t *testing.T) {
	length := 32
	str, err := GenerateRandomAlphabeticString(length)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(str) != length {
		t.Errorf("expected length %d, got %d", length, len(str))
	}
}

func TestGenerateRandomAlphabeticString_IsAlphabetic(t *testing.T) {
	str, err := GenerateRandomAlphabeticString(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, c := range str {
		if !unicode.IsLetter(c) {
			t.Errorf("unexpected non-letter character: %q", c)
		}
	}
}

func TestGenerateRandomString_CustomCharset(t *testing.T) {
	customCharset := "XYZ"
	str, err := GenerateRandomString(100, customCharset)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, c := range str {
		if c != 'X' && c != 'Y' && c != 'Z' {
			t.Errorf("character %q not in custom charset", c)
		}
	}
}

func TestGenerateRandomString_EmptyCharset(t *testing.T) {
	_, err := GenerateRandomString(10, "")
	if err == nil {
		t.Errorf("expected error for empty charset, got nil")
	}
}
