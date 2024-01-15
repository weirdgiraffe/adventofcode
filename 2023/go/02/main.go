package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pick struct {
	Red   int
	Green int
	Blue  int
}

func (p Pick) String() string {
	return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", p.Red, p.Green, p.Blue)
}

type Game struct {
	ID    int
	Picks []Pick
}

func (g Game) String() string {
	s := fmt.Sprintf("Game %d:\n", g.ID)
	for i := range g.Picks {
		if i != 0 {
			s += "\n"
		}
		s += "\t" + g.Picks[i].String()
	}
	return s
}

func (g Game) Possible(red int, green int, blue int) bool {
	for _, p := range g.Picks {
		if p.Red > red || p.Green > green || p.Blue > blue {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	var red, green, blue int
	for _, p := range g.Picks {
		red = max(red, p.Red)
		green = max(green, p.Green)
		blue = max(blue, p.Blue)
	}
	return red * green * blue
}

func parseDice(pick string) (Pick, error) {
	l := strings.Split(pick, ",")
	var p Pick
	for i := range l {
		ll := strings.Split(strings.TrimSpace(l[i]), " ")
		i, err := strconv.Atoi(ll[0])
		if err != nil {
			return p, fmt.Errorf("failed to number: %w", err)
		}

		switch ll[1] {
		case "red":
			p.Red = i
		case "green":
			p.Green = i
		case "blue":
			p.Blue = i
		default:
			return p, fmt.Errorf("wrong color: %s", ll[1])
		}
	}
	return p, nil
}

func parsePics(picks string) ([]Pick, error) {
	l := strings.Split(picks, ";")
	if len(l) == 0 {
		return nil, fmt.Errorf("empty picks")
	}
	var err error
	p := make([]Pick, len(l))
	for i := range l {
		p[i], err = parseDice(l[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse dice %q: %w", l[i], err)
		}
	}
	return p, nil
}

func parseLine(line string) Game {
	l := strings.Split(line, ":")
	if len(l) != 2 {
		panic("wrong line format: " + line)
	}

	id, err := strconv.Atoi(strings.TrimPrefix(l[0], "Game "))
	if err != nil {
		panic("wrong game id format: " + err.Error() + ":" + line)
	}
	picks, err := parsePics(l[1])
	if err != nil {
		panic("failed to parse picks: " + err.Error() + ":" + line)
	}
	return Game{
		ID:    id,
		Picks: picks,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		game := parseLine(scanner.Text())
		power := game.Power()
		sum += power
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read line: %v", err)
		os.Exit(1)
	}
	fmt.Println("sum=", sum)
}
