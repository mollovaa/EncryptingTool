package transposition_test

import "testing"
import transposition "../transposition"

func TestEncryptWithPositiveKey(t *testing.T) {
	result := transposition.Encrypt("security", 4)
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	result := transposition.Encrypt("security", -4)
	if result != "steiycru" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "steiycru")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	result := transposition.Decrypt("steiycru", 4)
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	result := transposition.Decrypt("steiycru", -4)
	if result != "security" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "security")
	}
}
