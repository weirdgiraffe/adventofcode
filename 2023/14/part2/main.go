package main

import (
	"bufio"
	"fmt"
	"os"
)

func totheright(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	b := make([]byte, len(lines))
	out := make([]string, len(lines[0]))

	for i := 0; i < len(out); i++ {
		for j, line := range lines {
			b[len(b)-j-1] = line[i]
		}
		out[i] = string(b)
	}
	return out
}

func totheleft(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	b := make([]byte, len(lines))
	out := make([]string, len(lines[0]))

	for i := 0; i < len(out); i++ {
		for j, line := range lines {
			b[j] = line[len(line)-i-1]
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
			// fmt.Printf("i=%d j=%d s=%s\n", i, j, string(b))
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
	fmt.Println("original")
	for i, s := range pattern {
		fmt.Printf("%2d | %s\n", i, s)
	}

	l := pattern
	for i := 0; i < 1; i++ {
		l = totheright(tilt(totheleft(l)))
		fmt.Println("north")
		for i, s := range l {
			fmt.Printf("%2d | %s\n", i, s)
		}
		l = tilt(l)
		fmt.Println("east")
		for i, s := range l {
			fmt.Printf("%2d | %s\n", i, s)
		}

		l = totheleft(tilt(totheright(l)))
		fmt.Println("south")
		for i, s := range l {
			fmt.Printf("%2d | %s\n", i, s)
		}

		l = totheright(totheright(tilt(totheleft(totheleft(l)))))
		fmt.Println("west")
		for i, s := range l {
			fmt.Printf("%2d | %s\n", i, s)
		}
	}

	return load(l)
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
