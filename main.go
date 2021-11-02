package main

import (
	"fmt"
	"os"
	"rlp-decoder/decode"
)

func main() {

	if len(os.Args) <= 1 {
		panic("Please provide an rlp encoded message as an argument")
	}

	// retrieve string from args
	interceptedMessage := os.Args[1]

	//check if its a valid hex string
	if !decode.IsValidHexString(interceptedMessage) {
		panic("Not a valid hex string")
	}

	//decode retrieved string to byte slice
	decodedBytes := decode.StrToByteSlice(interceptedMessage)

	// Get rlp decoded plain text, both as string and slice
	resultString, resultList := decode.Decode(decodedBytes)

	if resultString != "" && len(resultList) == 0 {
		// If decoded data is a string
		fmt.Println(resultString)
	} else if len(resultList) > 0 {
		// If decoded data is a list
		fmt.Println(resultList)
	} else if resultString == "" {
		fmt.Println("could not decode message")
	}
}
