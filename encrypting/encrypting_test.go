package encrypting

import "testing"
import "ceaser"

func TestDefineCipherChoiceKeyIsValid(t *testing.T) {
	ceaserCipher := &ceaser.Ceaser{}
	var ciphersMap = map[int]Cipher{
		1: ceaserCipher}

	result, _ := DefineCipherChoice(ciphersMap, "1")
	if result.Name() != ceaserCipher.Name() {
		t.Errorf("Define cipher text is incorrect, got %s, wanted %s", result, "Ceaser")
	}
}

func TestDefineCipherChoiceKeyIsInvalid(t *testing.T) {
	ceaserCipher := &ceaser.Ceaser{}
	var ciphersMap = map[int]Cipher{
		1: ceaserCipher}

	_, err := DefineCipherChoice(ciphersMap, "5")
	if err == nil {
		t.Errorf("Define cipher text is incorrect, got result, wanted error")
	}
}

func TestDefineOperationChoiceKeyIsValid(t *testing.T) {
	result, _ := DefineOperationChoice("1")

	if result != Encrypt {
		t.Errorf("Define operation choice is incorrect, got %s, wanted %s", result, "Encrypt")
	}
}

func TestDefineOperationChoiceKeyIsInvalid(t *testing.T) {
	_, err := DefineOperationChoice("7")

	if err == nil {
		t.Errorf("Define operation choice is incorrect, got result, wanted error")
	}
}
