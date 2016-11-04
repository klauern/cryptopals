package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const four = "4.txt"

func TestFindSingleCharXor(t *testing.T) {
	f, err := os.Open(four)
	if err != nil {
		t.Fatalf("Unable to open file %s, Error %v", four, err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	score, all, err := DetectSingleCharXor(lines)
	if err != nil {
		t.Fatalf("Error getting best score: %v", err)
	}
	fmt.Printf("Best Score: %d\nBest Score's Line: %s\nBest from Each Line: \n", score.bestScore, score.best)
	for _, v := range all {
		fmt.Println(v.best)
	}
}
