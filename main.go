package main

import (
	"fmt"
	"os"
	"reflect"
	"rlp-decoder/decode"
)

/*
RLP Decoding
1. Obtain byte string and check length for even vs odd. odd->invalid hex string
2. Convert given string to byte slice.
3. Iterate over byte slice:
- [00 ... 127] => data of type string and should be decoded as it is
- [128 ... 183] => data is a short string - defines type
- [184 ... 191] => data is long string - defines type
- [192 ... 247] => data is a short list - defines type
- [248 ... 255] => data is a long list - defines type
4. If byte falls in the above ranges that defines the type of data, length can be found by (current_byte - first_byte_from_range).
5. Parse until end of byte slice
*/

func main() {

	defer os.Exit(0)

	// retrieve string from args
	interceptedMessage := os.Args[1:][0]

	//check if its a valid hex string
	if !decode.IsValidHexString(interceptedMessage) {
		panic("Not a valid hex string")
	}

	//decode retrieved string to byte slice
	decodedBytes := decode.StrToByteSlice(interceptedMessage)

	// iterate through byte slice
	for i := 0; i < len(decodedBytes); i++ {
		fmt.Println(decodedBytes[i], reflect.TypeOf(decodedBytes[i]))
	}

	fmt.Println(decodedBytes)

}
