package encrypting

import "fmt"
import "sort"
import "strconv"

// Custom type Cipher which contains all needed cipher's method.
type Cipher interface {
	IncreaseTimesUsed()
	GetTimesUsed() int
	Name() string
	Encrypt(string, string) string
	Decrypt(string, string) string
}

// Function which prints all available ciphers with their corresponding keys in the map
func PrintAvailableCiphers(ciphers map[int]Cipher) {
	var keys []int
	for key := range ciphers {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Available ciphers:")
	for _, key := range keys {
		fmt.Println(key, ":", ciphers[key].Name())
	}
}

// Validates if the input is a corresponding number to a cipher.
// If yes, the number is returned, else error is thrown.
func DefineCipherChoice(ciphers map[int]Cipher, choice string) (Cipher, error) {
	key, err := strconv.Atoi(choice)

	if err == nil && key >= 1 && key <= len(ciphers) {
		return ciphers[key], nil
	}
	return nil, fmt.Errorf("Invalid choice.")
}

// Custom type OperationEnum with 2 options: Encrypt, Decrypt
type OperationEnum int

const (
	Encrypt OperationEnum = 1
	Decrypt OperationEnum = 2
	Invalid OperationEnum = 0
)

// stringify function of the OperationEnum
func (e OperationEnum) String() string {
	switch e {
	case Encrypt:
		return "Encryption"
	case Decrypt:
		return "Decryption"
	default:
		return "InvalidOperation"
	}
}

// Validates if the input is a number corresponding to a operarion.
// If yes, the operationEnum is returned, else error is thrown.
func DefineOperationChoice(choice string) (OperationEnum, error) {
	option, err := strconv.Atoi(choice)

	if err == nil && (option == 1 || option == 2) {
		switch option {
		case 1:
			return Encrypt, nil
		case 2:
			return Decrypt, nil
		}
	}
	return Invalid, fmt.Errorf("Invalid operation.")
}
