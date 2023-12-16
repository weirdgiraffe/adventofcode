package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrangements(t *testing.T) {
	tt := []struct {
		In  string
		Exp int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
		{"?#???#?.#???? 6,1,1", 6},
		{"???.##??#??????#???# 1,13,1", 3},
		{".??#.?#??#????#?? 2,4,1,1,1", 2},
		{"???#??#.???#.?#??#?# 4,2,1,2,1,4", 1},
		{"????#?##???#.??.? 2,2,6,1", 3},
		{"???? 3", 2},
		{"???# 3", 1},
		{"#??? 3", 1},
		{"?#?? 3", 2},
		{"??#? 3", 2},
		{".??? 3", 1},
		{"???. 3", 1},
		{"???? 1,1", 3},
		{".??? 1,1", 1},
		{"?.?? 1,1", 2},
		{"??.? 1,1", 2},
		{"???. 1,1", 1},
		{"?.?. 1,1", 1},
	}
	for i, tc := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			n := arrangements(record(tc.In))
			require.Equal(t, tc.Exp, n)
		})
	}
}

/*
????#?##???#.??.? 2,2,6,1

##.##.######.??.? 2,2,6,1
	##.##.######....# 2,2,6,1
	##.##.######..#.. 2,2,6,1
	##.##.######.#... 2,2,6,1
*/
