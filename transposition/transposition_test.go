package transposition

import "testing"

func TestEncryptWithPositiveKey(t *testing.T) {
	result := Encrypt("security", 4)
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	result := Encrypt("security", -4)
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	result := Decrypt("steiycru", 4)
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	result := Decrypt("steiycru", -4)
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}
