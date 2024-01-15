package main

import (
	"bufio"
	"fmt"
	"os"
)

type tupple struct {
	d int
	n int
}

func stringdigit(s string) *int {
	if len(s) < 2 {
		return nil
	}
	next := func(suffix string, digit int) func(s string) *int {
		return func(s string) *int {
			if len(s) < len(suffix) {
				return nil
			}
			if s[:len(suffix)] == suffix {
				return &digit
			}
			return nil
		}
	}

	m := map[string]func(string) *int{
		"on": next("e", 1),
		"tw": next("o", 2),
		"th": next("ree", 3),
		"fo": next("ur", 4),
		"fi": next("ve", 5),
		"si": next("x", 6),
		"se": next("ven", 7),
		"ei": next("ght", 8),
		"ni": next("ne", 9),
	}
	if parse, ok := m[s[:2]]; ok {
		return parse(s[2:])
	}
	return nil
}

func digit(s string) *int {
	if len(s) == 0 {
		return nil
	}
	m := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}
	if digit, ok := m[s[0]]; ok {
		return &digit
	}
	return stringdigit(s)
}

func firstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if d := digit(s[i:]); d != nil {
			return *d
		}
	}
	panic("no first digit found: " + s)
}

func lastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if d := digit(s[i:]); d != nil {
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
