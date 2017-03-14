package main

import (
	challenge "github.com/klauern/cryptopals/set_1/challenge_3"
)

// StringCipherScore represents a list of scores for a given
// slice of possibilities
type StringCipherScore struct {
	line  string
	score int
}

func worker(lineIn <-chan string, resultOut chan<- *StringCipherScore) {
	for line := range lineIn {
		best := BestCipherFromString(line)
		resultOut <- best
	}
}

// DetectSingleCharXor will take a slice of string and find the one among it
// that contains a decipherable decoding using one Xor operation against
// a character
func DetectSingleCharXor(lines []string) (*StringCipherScore, []*StringCipherScore, error) {
	var scores []*StringCipherScore

	linesCh := make(chan string, 50)
	resultsCh := make(chan *StringCipherScore, 50)

	// create workers for processing scores
	for w := 1; w <= 5; w++ {
		go worker(linesCh, resultsCh)
	}

	// run a goroutine to pass lines of input to the lineCh channel
	go func(lines []string) {
		for _, line := range lines {
			linesCh <- line
		}
		close(linesCh)
	}(lines)

	best := &StringCipherScore{}
	for range lines {
		score := <-resultsCh
		scores = append(scores, score)
	}

	for _, score := range scores {
		if score.score > best.score {
			best = score
		}
	}
	return best, scores, nil
}

// BestCipherFromString return a *StringCipherScore, representing the best
// possible cipher decoding from a given string.
func BestCipherFromString(line string) *StringCipherScore {
	best := &StringCipherScore{}
	byteLine := []byte(line)

	for i := uint16(0); i <= 0x255; i++ {
		best.addCipher(rune(i), byteLine)
	}

	// best.AddCipherFromRuneRange(unicode.ASCII_Hex_Digit.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.Punct.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.White_Space.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.Letter.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.Common.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.Digit.R16, byteLine)
	// best.AddCipherFromRuneRange(unicode.Other_Alphabetic.R16, byteLine)

	// for _, r16 := range unicode.ASCII_Hex_Digit.R16 {
	// 	for c := r16.Lo; c <= r16.Hi; c += r16.Stride {
	// 		fmt.Printf("%v", rune(c))
	// 		best.addCipher(rune(c), byteLine)
	// 	}
	// }
	// for _, r32 := range unicode.ASCII_Hex_Digit.R32 {
	// for c := r32.Lo; c <= r32.Hi; c += r32.Stride {
	// fmt.Printf("%v", rune(c))
	// best.addCipher(rune(c), byteLine)
	// }
	// }
	// for _, l := range unicode.Letter.R16 {
	// for c := l.Lo; c <= l.Hi; c += l.Stride {
	// best.addCipher(rune(c), byteLine)
	// }
	// }
	// for _, l := range unicode.Letter.R32 {
	// 	for c := l.Lo; c <= l.Hi; c += l.Stride {
	// 		best.addCipher(rune(c), byteLine)
	// 	}
	// }
	return best
}

func (best *StringCipherScore) addCipher(c rune, line []byte) {
	str, score := challenge.ScoreCipher(c, line)
	if score > best.score {
		best.line = str
		best.score = score
	}
}
