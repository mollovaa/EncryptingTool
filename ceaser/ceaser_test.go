package ceaser

import "testing"

func TestEncryptWithPositiveKey(t *testing.T) {
	result := Encrypt("Hello", 3)
	if result != "Khoor" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Khoor")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	result := Encrypt("Welcome", -3)
	if result != "Tbizljb" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Tbizljb")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	result := Decrypt("Khoor", 3)
	if result != "Hello" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Hello")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	result := Decrypt("Tbizljb", -3)
	if result != "Welcome" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Welcome")
	}
}
