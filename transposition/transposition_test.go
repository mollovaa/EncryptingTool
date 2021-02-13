package transposition

import "testing"

func TestEncryptWithPositiveKey(t *testing.T) {
	tr := &Transposition{}
	result := tr.Encrypt("security", "4")
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	tr := &Transposition{}
	result := tr.Encrypt("security", "-4")
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	tr := &Transposition{}
	result := tr.Decrypt("steiycru", "4")
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	tr := &Transposition{}
	result := tr.Decrypt("steiycru", "-4")
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}
