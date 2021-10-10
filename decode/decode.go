package decode

import "encoding/hex"

const plainString = "plain_string"
const shortString = "type_short_string"
const longString = "type_long_string"
const shortList = "type_short_list"
const longList = "type_long_list"

func isValidHexString(str string) bool {
	return len(str)%2 == 0
}

func strToByteSlice(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}

func getType(val uint8) string {
	if val > 0 && val <= 127 {
		return plainString
	} else if val >= 128 && val <= 183 {
		return shortString
	} else if val >= 184 && val <= 191 {
		return longString
	} else if val >= 192 && val <= 247 {
		return shortList
	} else if val >= 248 && val <= 255 {
		return longList
	}

	return "Invalid byte!"
}
