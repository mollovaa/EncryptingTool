package monoalphabetic

import "testing"

func TestEncryptWithShorterKeyThenPlaintext(t *testing.T) {
	m := &MonoAlphabetic{}
	result := m.Encrypt("WELCOME", "Water")
	if result != "FRQTNPR" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "FRQTNPR")
	}
}
func TestEncryptWithSameLenghts(t *testing.T) {
	m := &MonoAlphabetic{}
	result := m.Encrypt("WELCOME", "testing")
	if result != "DNUSPRN" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "DNUSPRN")
	}
}

func TestDecryptWithShorterKeyThenPlaintext(t *testing.T) {
	m := &MonoAlphabetic{}
	result := m.Decrypt("DNUSPRN", "testing")
	if result != "WELCOME" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "WELCOME")
	}
}

func TestDecryptWithSameLengths(t *testing.T) {
	m := &MonoAlphabetic{}
	result := m.Decrypt("FRQTNPR", "Water")
	if result != "WELCOME" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "WELCOME")
	}
}
