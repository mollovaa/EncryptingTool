package route

import "testing"

func TestEncryptWithPositiveKey(t *testing.T) {
	r := &Route{}
	result := r.Encrypt("Ceaser", "3")
	if result != "Cserae" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Cserae")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	r := &Route{}
	result := r.Encrypt("Ceaser", "-3")
	if result != "raeCse" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "raeCse")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	r := &Route{}
	result := r.Decrypt("Cserae", "3")
	if result != "Ceaser" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Ceaser")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	r := &Route{}
	result := r.Decrypt("raeCse", "-3")
	if result != "Ceaser" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Ceaser")
	}
}
