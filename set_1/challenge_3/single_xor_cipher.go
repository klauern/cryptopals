package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/sajari/fuzzy"
)

var model *fuzzy.Model

func init() {
	//sampleWords = fuzzy.SampleEnglish()
	model = fuzzy.NewModel()
	model.SetThreshold(1)
	model.SetDepth(2)
	mustLoadDictionary()
}

func mustLoadDictionary() {
	file, err := os.Open("my.dict")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	model.Train(words)
}

// ScoreCipher will produce an overall score of the cipher decryption.  It does this using three checks:
// 1. Whether the XOR produced a letter
// 2. Whether the XOR produced a Space character
// 3. Whether the entire phrase produced any properly spelled words
// Each pass of the options above produces a total score, with the sum being the total score of the
// cipher operation.
func ScoreCipher(char rune, encoding []byte) (string, int) {
	dest := make([]byte, len(encoding))
	total := 0
	for i, v := range encoding {
		dest[i] = v ^ byte(char)
		r := rune(dest[i])
		if unicode.In(r, unicode.Letter, unicode.Space) {
			total++
		} else if unicode.IsControl(r) || unicode.IsPunct(r) {
			total--
		}
	}
	output := fmt.Sprintf("%s", dest)
	total += ScoreWords(output)
	//fmt.Println(output)
	return fmt.Sprintf("%v", output), total
}

// ScoreWords will produce a score on a given phrase string based on how many
// spellable words are found within it.  This total is then returned to the
// caller.
func ScoreWords(phrase string) int {
	words := strings.Fields(phrase)
	score := 0
	for _, word := range words {
		//model.TrainWord(word)
		//fmt.Printf("Spell Check for %s is %s\n", word, model.SpellCheck(word))
		if word == model.SpellCheck(word) {
			score++
		}
	}
	return score * 100
}
