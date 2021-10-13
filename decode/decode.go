package decode

import (
	"encoding/hex"
	"fmt"
)

const PLAIN_STRING = "plain_string"
const SHORT_STRING = "type_short_string"
const LONG_STRING = "type_long_string"
const SHORT_LIST = "type_short_list"
const LONG_LIST = "type_long_list"

func init() {
	fmt.Println("Initializing decoder package...")
}

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
	if val > 0 && val <= 127 {
		return PLAIN_STRING
	} else if val >= 128 && val <= 183 {
		return SHORT_STRING
	} else if val >= 184 && val <= 191 {
		return LONG_STRING
	} else if val >= 192 && val <= 247 {
		return SHORT_LIST
	} else if val >= 248 && val <= 255 {
		return LONG_LIST
	}

	return "Invalid byte!"
}
