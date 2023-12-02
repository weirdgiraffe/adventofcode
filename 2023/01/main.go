package main

import (
	"bufio"
	"fmt"
	"os"
)

func digit(c byte) *int {
	m := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}
	if i, ok := m[c]; ok {
		return &i
	}
	return nil
}

func firstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if d := digit(c); d != nil {
			return *d
		}
	}
	panic("no first digit found: " + s)
}

func lastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if d := digit(c); d != nil {
			return *d
		}
	}
	panic("no last digit found: " + s)
}

func number(s string) int {
	first := firstDigit(s)
	last := lastDigit(s)
	return first*10 + last
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	line := 0
	for scanner.Scan() {
		line += 1
		sum += number(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read line: %v", err)
		os.Exit(1)
	}
	fmt.Println("sum=", sum)
}
