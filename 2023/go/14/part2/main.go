package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
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

// we expect the pattern already pointing to the noth here
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

func hash(l []string) string {
	h := sha256.New()
	for _, s := range l {
		h.Write([]byte(s))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func solve(pattern []string) int {
	m := make(map[string]int)
	l := totheleft(pattern)
	n := 1000000000
	cycle := false
	for i := 0; i < n; i++ {
		h := hash(l)
		if prev, ok := m[h]; ok {
			if !cycle {
				// we have detected a cycle. now we need to understand
				// how many iterations we actually need more in order
				// to match the target
				offt := i - prev
				ileft := n - i
				n = i + (ileft % offt)
				cycle = true
			}
		} else {
			m[h] = i
		}

		tilt(l)
		l = totheright(l)
		tilt(l)
		l = totheright(l)
		tilt(l)
		l = totheright(l)
		tilt(l)
		l = totheright(l)

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
