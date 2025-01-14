package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		desc string
		args []int
		val  int
		res  int
	}{
		{
			desc: "Positive1",
			args: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			val:  1,
			res:  1,
		},
		{
			desc: "Positive2",
			args: []int{10, 12, 23, 34, 45, 65, 76, 87, 98, 109},
			val:  98,
			res:  8,
		},
		{
			desc: "BorderCase1",
			args: []int{10, 12, 23, 34, 45, 65, 76, 87, 98, 109, 111},
			val:  10,
			res:  0,
		},
		{
			desc: "BorderCase2",
			args: []int{10, 12, 23, 34, 45, 65, 76, 87, 98, 109, 111},
			val:  111,
			res:  10,
		},
		{
			desc: "MinusNumber",
			args: []int{10, 12, 23, 34, 45, 65, 76, 87, 98, 109, 111},
			val:  -124,
			res:  -1,
		},
		{
			desc: "TooBigNumber",
			args: []int{10, 12, 23, 34, 45, 65, 76, 87, 98, 109, 111},
			val:  236236,
			res:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := BinarySearch(tC.args, tC.val)
			require.Equal(t, tC.res, got)
		})
	}
}
