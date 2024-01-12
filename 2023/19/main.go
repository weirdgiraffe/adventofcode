package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Rating struct {
	X, M, A, S int
}

func (r Rating) String() string {
	return fmt.Sprintf("{x=%d,m=%d,a=%d,s=%d}", r.X, r.M, r.A, r.S)
}

func (r Rating) Sum() int {
	return r.X + r.M + r.A + r.S
}

func (r *Rating) Set(c byte, rg int) {
	switch c {
	case 'x':
		r.X = rg
	case 'm':
		r.M = rg
	case 'a':
		r.A = rg
	case 's':
		r.S = rg
	}
}

func RatingFromString(s string) Rating {
	i := strings.Index(s, "{")
	var r Rating
	for _, part := range strings.Split(s[i+1:len(s)-1], ",") {
		v, err := strconv.Atoi(part[2:])
		if err != nil {
			panic("failed to parse rating number")
		}
		switch part[0] {
		case 'x':
			r.X = v
		case 'm':
			r.M = v
		case 'a':
			r.A = v
		case 's':
			r.S = v
		default:
			panic("failed to parse rating letter")
		}
	}
	return r
}

type Workflow struct {
	l []Action
}

func (w *Workflow) Execute(r *Rating) *Workflow {
	var result *Workflow
	for _, action := range w.l {
		result = action(r)
		if result != nil {
			break
		}
	}
	return result
}

type System struct {
	m map[string]*Workflow
}

type Action func(r *Rating) *Workflow

func (s *System) NewAction(name string, g RatingGetter, c ConditionMatcher, dst string) Action {
	return Action(func(r *Rating) *Workflow {
		if c.MatchCondition(g.GetRating(r)) {
			// fmt.Printf("%s -> %s\n", name, dst)
			return s.m[dst]
		}
		return nil
	})
}

type RatingGetter interface {
	GetRating(r *Rating) int
}

type RatingGetterFunc func(r *Rating) int

func (f RatingGetterFunc) GetRating(r *Rating) int {
	return f(r)
}

func makeRatingGetter(c byte) RatingGetter {
	switch c {
	case 'x':
		return RatingGetterFunc(func(r *Rating) int { return r.X })
	case 'm':
		return RatingGetterFunc(func(r *Rating) int { return r.M })
	case 'a':
		return RatingGetterFunc(func(r *Rating) int { return r.A })
	case 's':
		return RatingGetterFunc(func(r *Rating) int { return r.S })
	default:
		panic("wrong rating letter")
	}
}

type ConditionMatcher interface {
	MatchCondition(int) bool
}

type ConditionMatcherFunc func(int) bool

func (f ConditionMatcherFunc) MatchCondition(v int) bool {
	return f(v)
}

func makeConditionMatcher(c byte, number int) ConditionMatcher {
	switch c {
	case '>':
		return ConditionMatcherFunc(func(v int) bool { return v > number })
	case '<':
		return ConditionMatcherFunc(func(v int) bool { return v < number })
	default:
		panic("wrong condition letter")
	}
}

var (
	accept = &Workflow{}
	reject = &Workflow{}
)

func ParseWorkflows(lines []string) *System {
	system := &System{
		m: map[string]*Workflow{
			"A": accept,
			"R": reject,
		},
	}
	for _, line := range lines {
		i := strings.Index(line, "{")
		w := &Workflow{}
		name := line[:i]
		system.m[name] = w

		for _, step := range strings.Split(line[i+1:len(line)-1], ",") {
			l := strings.Split(step, ":")
			if len(l) > 2 {
				panic("wrong format for workflow")
			}
			if len(l) == 1 {
				// just destination
				w.l = append(w.l, system.NewAction(name,
					RatingGetterFunc(func(_ *Rating) int { return 0 }),
					ConditionMatcherFunc(func(_ int) bool { return true }),
					l[0]))
				continue
			}
			rating := makeRatingGetter(l[0][0])
			n, err := strconv.ParseInt(l[0][2:], 10, 64)
			if err != nil {
				panic("wrong number")
			}
			condition := makeConditionMatcher(l[0][1], int(n))
			w.l = append(w.l, system.NewAction(name, rating, condition, l[1]))
		}
	}
	return system
}

func ParseRatings(lines []string) []Rating {
	var l []Rating
	for _, line := range lines {
		l = append(l, RatingFromString(line))
	}
	return l
}

func parse(lines []string) (*System, []Rating) {
	for i, line := range lines {
		if line == "" {
			system := ParseWorkflows(lines[:i])
			ratings := ParseRatings(lines[i+1:])
			return system, ratings
		}
	}
	panic("no empty line")
}

func part1(system *System, ratings []Rating) int {
	la := make([]Rating, 0)
	for i, r := range ratings {
		fmt.Printf("%d | %d\r", i+1, len(ratings))
		w := system.m["in"]
		for w != nil {
			w = w.Execute(&r)
			if w == accept {
				la = append(la, r)
			}
		}
	}
	var sum int
	for _, r := range la {
		sum += r.Sum()
	}
	return sum
}

func part2(system *System, _ []Rating) int {
	total := int64(0)

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			n := atomic.LoadInt64(&total)
			fmt.Printf("%d\r", n)
		}
	}()
	defer ticker.Stop()

	var sum int64
	var wg sync.WaitGroup
	for x := 0; x < 4000; x++ {
		for m := 0; m < 4000; m++ {
			wg.Add(1)
			go func(x, m int) {
				defer wg.Done()
				for a := 0; a < 4000; a++ {
					for s := 0; s < 4000; s++ {
						atomic.AddInt64(&total, 1)
						r := Rating{X: x + 1, M: m + 1, A: a + 1, S: s + 1}
						w := system.m["in"]
						for w != nil {
							w = w.Execute(&r)
							if w == accept {
								atomic.AddInt64(&sum, int64(r.Sum()))
							}
						}
					}
				}
			}(x, m)
		}
	}
	wg.Wait()
	return int(sum)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	t0 := time.Now()
	sum := part1(parse(lines))
	fmt.Println()
	fmt.Println("part1=", sum, "elapsed:", time.Since(t0))
}
