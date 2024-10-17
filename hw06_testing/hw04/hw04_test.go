package hw04

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteID(t *testing.T) {
	testCases := []struct {
		desc string
		id   uint32
		res  uint32
	}{
		{
			desc: "Positive",
			id:   1236437372,
			res:  1236437372,
		},
		{
			desc: "BorderCase1",
			id:   0,
			res:  0,
		},
		{
			desc: "BorderCase2",
			id:   4294967295,
			res:  4294967295,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{}
			a.WriteID(tC.id)
			require.Equal(t, tC.res, a.id)
		})
	}
}

func TestWriteTitleAuthor(t *testing.T) {
	testCases := []struct {
		desc        string
		titleAuthor string
		res         string
	}{
		{
			desc:        "Positive",
			titleAuthor: "Artur",
			res:         "Artur",
		},
		{
			desc:        "NegativeNumbers",
			titleAuthor: "12-/3",
			res:         "12-/3",
		},
		{
			desc:        "NegativeSymbols",
			titleAuthor: "!-/.,",
			res:         "!-/.,",
		},
		{
			desc:        "TitleWithSpace",
			titleAuthor: "Sunshine my Friend and Be Happy",
			res:         "Sunshine my Friend and Be Happy",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{}
			a.WriteTitle(tC.titleAuthor)
			require.Equal(t, tC.res, a.title)
			a.WriteAuthor(tC.titleAuthor)
			require.Equal(t, tC.res, a.author)
		})
	}
}

func TestWriteYearSize(t *testing.T) {
	testCases := []struct {
		desc     string
		yearSize uint16
		res      uint16
	}{
		{
			desc:     "Positive",
			yearSize: 123,
			res:      123,
		},
		{
			desc:     "BorderCase1",
			yearSize: 0,
			res:      0,
		},
		{
			desc:     "BorderCase2",
			yearSize: 65535,
			res:      65535,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{}
			a.WriteYear(tC.yearSize)
			require.Equal(t, tC.res, a.year)
			a.WriteSize(tC.yearSize)
			require.Equal(t, tC.res, a.size)
		})
	}
}

func TestWriteRate(t *testing.T) {
	testCases := []struct {
		desc string
		rate float32
		res  float32
	}{
		{
			desc: "Positive",
			rate: 1.23,
			res:  1.23,
		},
		{
			desc: "NegativeInt",
			rate: 100,
			res:  100,
		},
		{
			desc: "BorderCase1",
			rate: 0,
			res:  0,
		},
		{
			desc: "BorderCase2",
			rate: 3.40282346638528859811704183484516925440e+38,
			res:  3.40282346638528859811704183484516925440e+38,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{}
			a.WriteRate(tC.rate)
			require.Equal(t, tC.res, a.rate)
		})
	}
}

func TestGetID(t *testing.T) {
	testCases := []struct {
		desc string
		id   uint32
		res  uint32
	}{
		{
			desc: "Positive",
			id:   1236437372,
			res:  1236437372,
		},
		{
			desc: "BorderCase1",
			id:   0,
			res:  0,
		},
		{
			desc: "BorderCase2",
			id:   4294967295,
			res:  4294967295,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{id: tC.id}
			got := a.GetID()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestGetTitleAuthor(t *testing.T) {
	testCases := []struct {
		desc        string
		titleAuthor string
		res         string
	}{
		{
			desc:        "Positive",
			titleAuthor: "Artur",
			res:         "Artur",
		},
		{
			desc:        "NegativeNumbers",
			titleAuthor: "12-/3",
			res:         "12-/3",
		},
		{
			desc:        "NegativeSymbols",
			titleAuthor: "!-/.,",
			res:         "!-/.,",
		},
		{
			desc:        "TitleWithSpace",
			titleAuthor: "Sunshine my Friend and Be Happy",
			res:         "Sunshine my Friend and Be Happy",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{title: tC.titleAuthor, author: tC.titleAuthor}
			got := a.GetTitle()
			require.Equal(t, tC.res, got)
			got = a.GetAuthor()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestGetYearSize(t *testing.T) {
	testCases := []struct {
		desc     string
		yearSize uint16
		res      uint16
	}{
		{
			desc:     "Positive",
			yearSize: 123,
			res:      123,
		},
		{
			desc:     "BorderCase1",
			yearSize: 0,
			res:      0,
		},
		{
			desc:     "BorderCase2",
			yearSize: 65535,
			res:      65535,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{year: tC.yearSize, size: tC.yearSize}
			got := a.GetYear()
			require.Equal(t, tC.res, got)
			got = a.GetSize()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestGetRate(t *testing.T) {
	testCases := []struct {
		desc string
		rate float32
		res  float32
	}{
		{
			desc: "Positive",
			rate: 1.23,
			res:  1.23,
		},
		{
			desc: "NegativeInt",
			rate: 100,
			res:  100,
		},
		{
			desc: "BorderCase1",
			rate: 0,
			res:  0,
		},
		{
			desc: "BorderCase2",
			rate: 3.40282346638528859811704183484516925440e+38,
			res:  3.40282346638528859811704183484516925440e+38,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			a := &Book{rate: tC.rate}
			got := a.GetRate()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestCreateStruct(t *testing.T) {
	testCases := []struct {
		desc string
		x    uint8
		enum uint8
	}{
		{
			desc: "Positive",
			x:    size,
			enum: size,
		},
		{
			desc: "ZeroValue",
			x:    0,
			enum: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str := createStruct((tC.x))
			require.Equal(t, tC.enum, str.enum)
		})
	}
}

func TestComparator(t *testing.T) {
	testCases := []struct {
		desc  string
		enum  uint8
		book1 Book
		book2 Book
		res   bool
	}{
		{
			desc:  "PositiveTrue",
			enum:  0,
			book1: Book{year: 1990, size: 500, rate: 7.3},
			book2: Book{year: 1989, size: 400, rate: 3.3},
			res:   true,
		},
		{
			desc:  "PositiveFalse",
			enum:  2,
			book1: Book{year: 1990, size: 500, rate: 7.3},
			book2: Book{year: 1999, size: 700, rate: 12.3},
			res:   false,
		},
		{
			desc:  "ZeroValue",
			enum:  1,
			book1: Book{},
			book2: Book{},
			res:   false,
		},
		{
			desc:  "NotCorrectFieldToCompare",
			enum:  10,
			book1: Book{year: 1990, size: 500, rate: 7.3},
			book2: Book{year: 1999, size: 700, rate: 12.3},
			res:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str := CBooks{tC.enum}
			got := str.Comparator(tC.book1, tC.book2)
			if tC.res {
				require.True(t, tC.res, got)
			} else {
				require.False(t, tC.res, got)
			}
		})
	}
}
