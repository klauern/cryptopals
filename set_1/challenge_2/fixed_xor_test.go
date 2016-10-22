package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestFixedXor(t *testing.T) {
	input, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Fatalf("Expected to decode string, did not")
	}
	inXor, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		t.Fatalf("Expected to decode string, did not")
	}
	expected, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if err != nil {
		t.Fatalf("Expected to decode string, did not")
	}

	actual := FixedXor(input, inXor)
	fmt.Printf("%s\n", actual)
	comp := bytes.Compare([]byte(expected), actual)
	if comp != 0 {
		t.Fatalf("Expected %s, Got %s", expected, actual)
	}
}
