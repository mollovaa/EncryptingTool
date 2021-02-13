package ceaser

import "testing"

func TestEncryptWithPositiveKey(t *testing.T) {
	c := &Ceaser{}
	result := c.Encrypt("Hello", "3")
	if result != "Khoor" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Khoor")
	}
}

func TestEncryptWithNegativeKey(t *testing.T) {
	c := &Ceaser{}
	result := c.Encrypt("Welcome", "-3")
	if result != "Tbi`ljb" {
		t.Errorf("Encryption is incorrect, got %s, wanted %s", result, "Tbizljb")
	}
}

func TestDecryptWithPositiveKey(t *testing.T) {
	c := &Ceaser{}
	result := c.Decrypt("Khoor", "3")
	if result != "Hello" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Hello")
	}
}

func TestDecryptWithNegativeKey(t *testing.T) {
	c := &Ceaser{}
	result := c.Decrypt("Tbi`ljb", "-3")
	if result != "Welcome" {
		t.Errorf("Decryption is incorrect, got %s, wanted %s", result, "Welcome")
	}
}

func TestIncreaseTimesUsed(t *testing.T) {
	c := &Ceaser{0}
	c.IncreaseTimesUsed()
	result := c.GetTimesUsed()
	if result != 1 {
		t.Errorf("Increasing times used is incorrect, got %d, wanted 0.", result)
	}
}
