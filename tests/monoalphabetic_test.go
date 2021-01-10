package monoalphabetic_test

import "testing"
import monoalphabetic "../monoalphabetic"

func TestEncryptWithShorterKeyThenPlaintext(t *testing.T) {
	result := monoalphabetic.Encrypt("WELCOME", "Water")
	if result != "FRQTNPR" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "FRQTNPR")
	}
}
func TestEncryptWithSameLenghts(t *testing.T) {
	result := monoalphabetic.Encrypt("WELCOME", "testing")
	if result != "DNUSPRN" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "DNUSPRN")
	}
}

func TestDecryptWithShorterKeyThenPlaintext(t *testing.T) {
	result := monoalphabetic.Decrypt("DNUSPRN", "testing")
	if result != "WELCOME" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "WELCOME")
	}
}

func TestDecryptWithSameLengths(t *testing.T) {
	result := monoalphabetic.Decrypt("FRQTNPR", "Water")
	if result != "WELCOME" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "WELCOME")
	}
}
