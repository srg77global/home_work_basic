package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		ok      bool
		logFile string
		level   string
		output  string
	)
	const empty = "empty"

	flag.StringVar(&logFile, "file", empty, "log file for analysis")
	flag.StringVar(&level, "level", `"info"`, "log level for analysis")
	flag.StringVar(&output, "output", empty, "path of file with the statistic")
	flag.Parse()

	if logFile == empty {
		logFile, ok = os.LookupEnv("LOG_ANALYZER_FILE")
		if !ok {
			fmt.Println("uncorrected path for the log file")
			return
		}
	}
	if level == `"info"` {
		level, ok = os.LookupEnv("LOG_ANALYZER_LEVEL")
		if !ok {
			fmt.Println(`using the standart log level "info"`)
		}
	}
	if output == empty {
		output, ok = os.LookupEnv("LOG_ANALYZER_OUTPUT")
		if !ok {
			output = empty
		}
	}

	file, err := os.Open(logFile)
	if err != nil {
		fmt.Printf("error os.Open: %v\n", err)
		defer file.Close()
		return
	}

	buf, err := csvRead(file)
	if err != nil {
		fmt.Println(err)
	}

	result := levelChoose(level, buf, logFile)

	file.Close()

	if output == empty {
		fmt.Println(result)
		return
	}

	fileWritten, err := os.Create(output)
	if err != nil {
		fmt.Printf("error os.Create: %v", err)
	}

	_, err = fileWritten.WriteString(result)
	if err != nil {
		fmt.Printf("error fileWritten.WriteString: %v", err)
	}

	fileWritten.Close()
}

func csvRead(file *os.File) ([][]string, error) {
	csvFile := csv.NewReader(file)
	buffer, err := csvFile.ReadAll()
	if err != nil {
		err = fmt.Errorf("error csvFile.ReadAll: %w", err)
		return nil, err
	}
	return buffer, nil
}

func levelChoose(level string, buffer [][]string, logFile string) string {
	switch level {
	case `"info"`, "":
		lines, errors := 0, 0

		for _, args := range buffer {
			lines++
			if args[1] == " error" {
				errors++
			}
		}

		res := fmt.Sprintf("the log file contains %d lines and %d errors", lines, errors)
		return res
	case "january", "1":
		return csvMonth("1", buffer, logFile)
	case "february", "2":
		return csvMonth("2", buffer, logFile)
	case "march", "3":
		return csvMonth("3", buffer, logFile)
	case "april", "4":
		return csvMonth("4", buffer, logFile)
	case "may", "5":
		return csvMonth("5", buffer, logFile)
	case "june", "6":
		return csvMonth("6", buffer, logFile)
	case "july", "7":
		return csvMonth("7", buffer, logFile)
	case "august", "8":
		return csvMonth("8", buffer, logFile)
	case "september", "9":
		return csvMonth("9", buffer, logFile)
	case "october", "10":
		return csvMonth("10", buffer, logFile)
	case "november", "11":
		return csvMonth("11", buffer, logFile)
	case "december", "12":
		return csvMonth("12", buffer, logFile)
	}
	return "month is undefined"
}

func csvMonth(str string, buffer [][]string, logFile string) string {
	lines, errors, date := 0, 0, ""

	if str == "10" || str == "11" || str == "12" {
		date = str
	} else {
		dateSlice := make([]string, 0)
		dateSlice = append(dateSlice, "0")
		dateSlice = append(dateSlice, str)
		date = strings.Join(dateSlice, "")
	}

	if logFile == "log_2024.csv" {
		date += ".2024"
	} else if logFile == "log_2023.csv" {
		date += ".2023"
	}

	for _, args := range buffer {
		if args[0] == date {
			lines++
			if args[1] == " error" {
				errors++
			}
		}
	}

	res := fmt.Sprintf("the log file contains %d lines and %d errors", lines, errors)
	return res
}
