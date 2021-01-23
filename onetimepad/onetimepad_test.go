package onetimepad

import "testing"

func TestEncrypt(t *testing.T) {
	result := Encrypt("package", "letters")
	if result != "AEVDEXW" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "AEVDEXW")
	}
}

func TestDecrypt(t *testing.T) {
	result := Decrypt("AEVDEXW", "letters")
	if result != "PACKAGE" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "PACKAGE")
	}
}

