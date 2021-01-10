package ceaser_test

import "testing"
import ceaser "../ceaser"

func TestEncryptWithPositiveKey(t *testing.T) {
	result := ceaser.Encrypt("Hello", 3)
	if result != "Khoor" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Khoor")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	result := ceaser.Encrypt("Welcome", -3)
	if result != "Tbizljb" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Tbizljb")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	result := ceaser.Decrypt("Khoor", 3)
	if result != "Hello" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Hello")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	result := ceaser.Decrypt("Tbizljb", -3)
	if result != "Welcome" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Welcome")
	}
}
