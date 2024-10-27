package hw03

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateChessboard(t *testing.T) {
	testCases := []struct {
		desc   string
		size   int
		output string
	}{
		{
			desc:   "PositiveEven",
			size:   6,
			output: "# # # \n # # #\n# # # \n # # #\n# # # \n # # #\n",
		},
		{
			desc:   "PositiveUneven",
			size:   7,
			output: "# # # #\n # # # \n# # # #\n # # # \n# # # #\n # # # \n# # # #\n",
		},
		{
			desc:   "ZeroValue",
			size:   0,
			output: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := createChessboard(tC.size)
			require.Equal(t, tC.output, got)
		})
	}
}
