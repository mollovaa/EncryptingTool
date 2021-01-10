package route_test

import "testing"
import route "../route"

func TestEncryptWithPositiveKey(t *testing.T) {
	result := route.Encrypt("Ceaser", 3)
	if result != "Cserae" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Cserae")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	result := route.Encrypt("Ceaser", -3)
	if result != "raeCse" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "raeCse")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	result := route.Decrypt("Cserae", 3)
	if result != "Ceaser" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Ceaser")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	result := route.Decrypt("raeCse", -3)
	if result != "Ceaser" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Ceaser")
	}
}
