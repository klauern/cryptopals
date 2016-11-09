package challenge

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestSingleXorCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inBytes, err := hex.DecodeString(input)
	if err != nil {
		t.Fatalf("error decoding input string, %v", err)
	}
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxScore := -1
	maxStr := ""
	for _, v := range alphabet {
		str, score := ScoreCipher(v, inBytes)
		if score > maxScore {
			maxScore = score
			maxStr = str
		}
		fmt.Printf("input xor is %s, output is %s with score %d\n", strconv.QuoteRune(v), str, score)
	}
	fmt.Printf("Max Score is %d, with string \"%s\"", maxScore, maxStr)
}
