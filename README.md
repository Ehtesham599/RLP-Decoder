# RLP-Decoder
> An implementation of Recursive Length Prefix Decoder in Golang.

## RLP decoding

Ethereum uses Recursive Length Prefix to serialize/deserialize data.

RLP decoding has a specified byte range set in order to decipher endcoded data.
```
- [00 .. 7f]: Data is of type String and should be decoded as it is.
- [80 .. b7]: String and its a short string.
- [b8 .. bf]: String and its a long string.
- [c0 .. f7]: List and short list.
- [f8 .. ff]: List and it’s a long list.
```

Recursive length prefix decoding involves the following steps:
1. Retreive the first byte.
2. Check if byte is of type string. If yes, directly decipher.
3. Else, check if the byte falls in any of the specified range set.
4. Calculate the length by subtracting retrieved first byte from the first byte of the byte range set it falls into.
5. Parse until end of string

## Execution

To run the application, run `go run main.go <RLP_ENCODED_MESSAGE>`, passing a valid RLP encoded message as an argument. 
