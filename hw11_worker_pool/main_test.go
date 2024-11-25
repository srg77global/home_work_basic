package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		v    int
		res  int
		i    int
	}{
		{
			desc: "Positive1",
			v:    0,
			res:  1,
			i:    0,
		},
		{
			desc: "Positive2",
			v:    58923,
			res:  58924,
			i:    0,
		},
		{
			desc: "MinusNums",
			v:    -24781,
			res:  -24780,
			i:    0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			wg.Add(3)
			Counter(&tC.v, tC.i)
			require.Equal(t, tC.res, tC.v)
		})
	}
}
