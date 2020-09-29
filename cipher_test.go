package main

import (
	"testing"
)

func TestCipher(t *testing.T) {
	pass := "BACKENDARCHITECTURE"
	startChar := 'A'
	endChar := 'Z'
	cipher := cipher{pass, startChar, endChar}

	original := "GO IS AWESOME, AND REALLY FUN TO PLAY WITH"
	expectedEncoded := "HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"

	encoded := cipher.Encode(original)
	decoded := cipher.Decode(expectedEncoded)

	if encoded != expectedEncoded {
		t.Errorf("Failed to encode, got: %s, expected: %s", encoded, expectedEncoded)
	}

	if decoded != original {
		t.Errorf("Failed to decode, got: %s, expected: %s", decoded, original)
	}
}
