package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	seq := "HASH"
	n := solve(seq)
	require.Equal(t, 52, n)
}

func TestSeq(t *testing.T) {
	seq := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	n := solve(seq)
	require.Equal(t, 1320, n)
}
