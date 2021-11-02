package decode

import (
	"encoding/hex"
	"fmt"
	"os"
)

type DataType struct {
	Start uint8
	End   uint8
}

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
	plainStringType = DataType{Start: 0, End: 127}
	shortStringType = DataType{Start: 128, End: 183}
	longStringType  = DataType{Start: 184, End: 191}
	shortListType   = DataType{Start: 192, End: 247}
	longListType    = DataType{Start: 248, End: 255}
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
	if val > plainStringType.Start && val <= plainStringType.End {
		return PLAIN_STRING
	} else if val >= shortStringType.Start && val <= shortStringType.End {
		return SHORT_STRING
	} else if val >= longStringType.Start && val <= longStringType.End {
		return LONG_STRING
	} else if val >= shortListType.Start && val <= shortListType.End {
		return SHORT_LIST
	} else if val >= longListType.Start && val <= longListType.End {
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
			l := data[i] - shortStringType.Start

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += "\n" + res + "\n"
			decodedList = append(decodedList, decodedMessage)

			// increment index to after current type
			i += int(l + 1)

		case LONG_STRING:
			//Get length of string
			l := data[i] - longStringType.Start

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += "\n" + res + "\n"
			decodedList = append(decodedList, decodedMessage)

			// increment index to after current type
			i += int(l + 1)

		case SHORT_LIST:
			//Get length of list
			l := data[i] - shortListType.Start

			// pass values ranging in the indexes obtained from length recursivley
			res, _ := Decode(data[i+1 : i+int(l+1)])

			// Append deciphered message
			decodedMessage += res
			decodedList = append(decodedList, res)

			// increment index to after current type
			i += int(l + 1)

		case LONG_LIST:
			//Get length of list
			l := data[i] - longListType.Start

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
