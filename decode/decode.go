package decode

import (
	"encoding/hex"
	"fmt"
	"os"
)

const (
	PLAIN_STRING = "plain string"
	SHORT_STRING = "type short string"
	LONG_STRING  = "type long string"
	SHORT_LIST   = "type short list"
	LONG_LIST    = "type long list"
	INVALID_TYPE = "invalid type"
)

// Range that follows the rule to decipher rlp encoded messages
var (
	PLAIN_STRING_SLICE = []uint8{0, 127}
	SHORT_STRING_SLICE = []uint8{128, 183}
	LONG_STRING_SLICE  = []uint8{184, 191}
	SHORT_LIST_SLICE   = []uint8{192, 247}
	LONG_LIST_SLICE    = []uint8{248, 255}
)

// Check if input hex string is valid
func IsValidHexString(str string) bool {
	if str != "" {
		return len(str)%2 == 0
	}

	return false
}

// decode string to hex
func StrToByteSlice(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}

// Check range of decoded and return its type
func GetType(val uint8) string {
	if val > PLAIN_STRING_SLICE[0] && val <= PLAIN_STRING_SLICE[1] {
		return PLAIN_STRING
	} else if val >= SHORT_STRING_SLICE[0] && val <= SHORT_STRING_SLICE[1] {
		return SHORT_STRING
	} else if val >= LONG_STRING_SLICE[0] && val <= LONG_STRING_SLICE[1] {
		return LONG_STRING
	} else if val >= SHORT_LIST_SLICE[0] && val <= SHORT_LIST_SLICE[1] {
		return SHORT_LIST
	} else if val >= LONG_LIST_SLICE[0] && val <= LONG_LIST_SLICE[1] {
		return LONG_LIST
	}

	return INVALID_TYPE
}

// Decodes uint8 slice recursively by following the recursive length prefix method
func Decode(data []uint8) (string, []string) {

	var decodedList []string
	var decodedMessage string

	for i := 0; i < len(data); i++ {

		switch input_type := GetType(data[i]); input_type {
		//Base case
		case PLAIN_STRING:
			// decipher string directly as it is
			decodedMessage += string(data[i])

		case SHORT_STRING:
			//Get length of string
			l := data[i] - SHORT_STRING_SLICE[0]

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += "\n" + res + "\n"
			decodedList = append(decodedList, decodedMessage)

			// increment index to after current type
			i += int(l + 1)

		case LONG_STRING:
			//Get length of string
			l := data[i] - LONG_STRING_SLICE[0]

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += "\n" + res + "\n"
			decodedList = append(decodedList, decodedMessage)

			// increment index to after current type
			i += int(l + 1)

		case SHORT_LIST:
			//Get length of list
			l := data[i] - SHORT_LIST_SLICE[0]

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += res
			decodedList = append(decodedList, res)

			// increment index to after current type
			i += int(l + 1)

		case LONG_LIST:
			//Get length of list
			l := data[i] - LONG_LIST_SLICE[0]

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += res
			decodedList = append(decodedList, res)

			// increment index to after current type
			i += int(l + 1)

		case INVALID_TYPE:
			fmt.Println("Invalid RLP code!")
			os.Exit(1)
		}
	}
	return decodedMessage, decodedList
}
