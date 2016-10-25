package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const file = "4.txt"

func TestFindSingleCharXor(t *testing.T) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("Unable to open file %s, Error %v", file, err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	score, err := DetectSingleCharXor(lines)
	if err != nil {
		t.Fatalf("Error getting best score: %v", err)
	}
	fmt.Printf("Best Score: %d\nLines: %v", score.bestScore, score.possibilities)
}
