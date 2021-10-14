package decode

import (
	"encoding/hex"
)

const (
	PLAIN_STRING = "plain string"
	SHORT_STRING = "type short string"
	LONG_STRING  = "type long string"
	SHORT_LIST   = "type short list"
	LONG_LIST    = "type long list"
	INVALID_BYTE = "invalid byte"
)

var (
	PLAIN_STRING_SLICE = []uint8{0, 127}
	SHORT_STRING_SLICE = []uint8{128, 183}
	LONG_STRING_SLICE  = []uint8{184, 191}
	SHORT_LIST_SLICE   = []uint8{192, 247}
	LONG_LIST_SLICE    = []uint8{248, 255}
)

func IsValidHexString(str string) bool {
	return len(str)%2 == 0
}

func StrToByteSlice(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}

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

	return "Invalid input!"
}

func Decode(data []uint8) (string, []string) {

	var decodedList []string
	var decodedMessage string
	for i := 0; i < len(data); i++ {
		switch byte_type := GetType(data[i]); byte_type {
		//Base case
		case PLAIN_STRING:
			// decode string directly as it is
			decodedMessage += string(data[i])
		case SHORT_STRING:
			//Get length of string
			l := data[i] - SHORT_STRING_SLICE[0]
			res, _ := Decode(data[i+1 : i+int(l+1)])
			decodedMessage += res + " "
			decodedList = append(decodedList, decodedMessage)
			i += int(l + 1)
		case LONG_STRING:
			//Get length of string
			l := data[i] - LONG_STRING_SLICE[0]
			res, _ := Decode(data[i+1 : i+int(l+1)])
			decodedMessage += res + " "
			decodedList = append(decodedList, decodedMessage)
			i += int(l + 1)
		case SHORT_LIST:
			//Get length list
			l := data[i] - SHORT_LIST_SLICE[0]
			res, _ := Decode(data[i+1 : i+int(l+1)])
			decodedMessage += res
			decodedList = append(decodedList, res)
			i += int(l + 1)
		case LONG_LIST:
			//Get length list
			l := data[i] - LONG_LIST_SLICE[0]
			res, _ := Decode(data[i+1 : i+int(l+1)])
			decodedMessage += res
			decodedList = append(decodedList, res)
			i += int(l + 1)
		default:
			decodedMessage += ""
		}
	}
	return decodedMessage, decodedList
}
