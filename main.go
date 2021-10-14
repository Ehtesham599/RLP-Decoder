package main

import (
	"fmt"
	"os"
	"rlp-decoder/decode"
)

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
	// for i := 0; i < len(decodedBytes); i++ {
	// 	fmt.Println(decodedBytes[i], reflect.TypeOf(decodedBytes[i]), string(decodedBytes[i]), decode.GetType(decodedBytes[i]))
	// }
	resString, resList := decode.Decode(decodedBytes)
	if resString != "" && len(resList) == 0 {
		fmt.Println(resString)
	} else if len(resList) > 0 {
		fmt.Println(resList)
	}
}
