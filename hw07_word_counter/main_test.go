package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		desc string
		str  string
		res  map[string]int
	}{
		{
			desc: "Positive",
			str:  "Hello, Hello-hello World!",
			res:  map[string]int{"hello": 3, "world": 1},
		},
		{
			desc: "EmptyValue",
			str:  "",
			res:  map[string]int{},
		},
		{
			desc: "OnlySymbols",
			str:  "!? !., ,",
			res:  map[string]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countWords(tC.str)
			require.Equal(t, tC.res, got)
		})
	}
}
