package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		bookVar *SBook
		bookU   *SBook
	}{
		{
			desc: "Positive",
			bookVar: &SBook{
				ID:     4125,
				Title:  "NewBook",
				Author: "Chkalov",
				Year:   1990,
				Size:   535,
				Rate:   6.7,
			},
			bookU: &SBook{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			j, err := json.Marshal(tC.bookVar)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(j, &tC.bookU)
			if err != nil {
				fmt.Println(err)
			}
			require.Equal(t, tC.bookVar, tC.bookU)
		})
	}
}
