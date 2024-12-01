package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCsvRead(t *testing.T) {
	testCases := []struct {
		desc    string
		logFile string
		res     [][]string
	}{
		{
			desc:    "Positive_log_2023.csv",
			logFile: "log_2023.csv",
			res: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
		},
		{
			desc:    "Negative_log_1990.csv",
			logFile: "log_1990.csv",
			res:     [][]string(nil),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			file, _ := os.Open(tC.logFile)
			buf, _ := csvRead(file)
			require.Equal(t, tC.res, buf)
		})
	}
}

func TestLevelChoose(t *testing.T) {
	testCases := []struct {
		desc    string
		level   string
		buf     [][]string
		logFile string
		res     string
	}{
		{
			desc:  `Positive_"info"_log_2023.csv`,
			level: `"info"`,
			buf: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
			logFile: "log_2023.csv",
			res:     "the log file contains 20 lines and 14 errors",
		},
		{
			desc:  `Positive_5_log_2023.csv`,
			level: "5",
			buf: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
			logFile: "log_2023.csv",
			res:     "the log file contains 1 lines and 0 errors",
		},
		{
			desc:  `Negative_0_log_2023.csv`,
			level: "0",
			buf: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
			logFile: "log_2023.csv",
			res:     "month is undefined",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := levelChoose(tC.level, tC.buf, tC.logFile)
			require.Equal(t, tC.res, got)
		})
	}
}

func TestCsvMonth(t *testing.T) {
	testCases := []struct {
		desc    string
		level   string
		buffer  [][]string
		logFile string
		res     string
	}{
		{
			desc:  `Positive_"6"_log_2023.csv`,
			level: "6",
			buffer: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
			logFile: "log_2023.csv",
			res:     "the log file contains 2 lines and 2 errors",
		},
		{
			desc:  `Negative_"0"_log_2023.csv`,
			level: "0",
			buffer: [][]string{
				{"01.2023", " error", " file not found"},
				{"02.2023", " ", " succeed opening"},
				{"02.2023", " error", " file is empty"},
				{"03.2023", " ", " succeed opening"},
				{"04.2023", " ", " succeed opening"},
				{"04.2023", " error", " file is empty"},
				{"04.2023", " error", " file not found"},
				{"05.2023", " ", " succeed opening"},
				{"06.2023", " error", " file not found"},
				{"06.2023", " error", " file is empty"},
				{"07.2023", " error", " file is empty"},
				{"08.2023", " error", " file is empty"},
				{"09.2023", " error", " file not found"},
				{"09.2023", " error", " file is empty"},
				{"09.2023", " ", " succeed opening"},
				{"10.2023", " error", " file not found"},
				{"11.2023", " error", " file is empty"},
				{"11.2023", " ", " succeed opening"},
				{"12.2023", " error", " file not found"},
				{"12.2023", " error", " file is empty"},
			},
			logFile: "log_2023.csv",
			res:     "the log file contains 0 lines and 0 errors",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := csvMonth(tC.level, tC.buffer, tC.logFile)
			require.Equal(t, tC.res, got)
		})
	}
}
