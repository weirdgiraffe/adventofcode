package main

import (
	"bufio"
	"fmt"
	"os"
)

func transpose(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}
	out := make([]string, len(lines[0]))
	b := make([]byte, len(lines))
	for i := 0; i < len(out); i++ {
		for j, line := range lines {
			b[j] = line[i]
		}
		out[i] = string(b)
	}
	return out

}

func load(l []string) int {
	n := len(l)
	sum := 0
	for _, s := range l {
		for j, c := range s {
			if c == 'O' {
				sum += (n - j)
			}
		}
	}
	return sum
}

func tilt(l []string) []string {
	for si := range l {
		b := []byte(l[si])

		var i, j = 0, 0
		for ; j < len(b) && b[j] != '.'; j++ {
		}

		for ; i < len(b); i++ {
			if b[i] == 'O' {
				// we have something to move
				if j < i && b[j] == '.' {
					b[j], b[i] = 'O', '.'
					j++
					for ; j < len(b) && b[j] == '#'; j++ {
					}
				}
			}
			if b[i] == '#' {
				for j = i; j < len(b) && b[j] == '#'; j++ {
				}
				for ; j < len(b) && b[j] == 'O'; j++ {
				}
			}
		}

		l[si] = string(b)
	}
	return l
}

func solve(pattern []string) int {
	tilted := tilt(transpose(pattern))
	return load(tilted)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	pattern := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		pattern = append(pattern, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read line: %v", err)
		os.Exit(1)
	}
	if len(pattern) > 0 {
		sum = solve(pattern)
	}
	fmt.Println("sum=", sum)
}
