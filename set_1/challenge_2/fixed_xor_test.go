package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestFixedXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	inXor := "686974207468652062756c6c277320657965"
	expected, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if err != nil {
		t.Fatalf("Expected to decode string, did not")
	}

	actual, err := FixedXorString(input, inXor)
	if err != nil {
		t.Fatalf("Error in parsing: %v", err)
	}
	fmt.Printf("%s\n", actual)
	comp := bytes.Compare([]byte(expected), actual)
	if comp != 0 {
		t.Fatalf("Expected %s, Got %s", expected, actual)
	}
}
