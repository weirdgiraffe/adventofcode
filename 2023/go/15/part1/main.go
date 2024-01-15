package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
	current := 0
	for _, c := range s {
		current += int(c)
		// fmt.Printf("1) %d\n", current)
		current *= 17
		// fmt.Printf("2) %d\n", current)
		current = current % 256
		// fmt.Printf("3) %d\n", current)
	}
	fmt.Printf("hash:%s:  %d\n", s, current)
	return current
}

func solve(seq string) int {
	sum := 0
	i, j := 0, 0
	for i = range seq {
		if seq[i] == ',' {
			sum += hash(seq[j:i])
			j = i + 1
		}
	}
	return sum + hash(seq[j:])
}

func main() {
	seq, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	sum := solve(strings.TrimSpace(string(seq)))
	fmt.Println("sum=", sum)
}
