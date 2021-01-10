package onetimepad_test

import "testing"
import 	onetimepad "../onetimepad"


func TestEncrypt(t *testing.T) {
	result := onetimepad.Encrypt("package", "letters")
	if result != "AEVDEXW" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "AEVDEXW")
	}
}

func TestDecrypt(t *testing.T) {
	result := onetimepad.Decrypt("AEVDEXW", "letters")
	if result != "PACKAGE" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "PACKAGE")
	}
}

