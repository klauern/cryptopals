package main

import "encoding/base64"
import "encoding/hex"

//var output = []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

//const hexencoding = ""

func ConvertHexToBase64(input string) string {
	inStr, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	output := base64.StdEncoding.EncodeToString(inStr)
	return output
}
