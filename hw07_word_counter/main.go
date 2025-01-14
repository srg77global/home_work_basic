package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(str string) map[string]int {
	strChanged := strings.ReplaceAll(str, ".", "")
	strChanged = strings.ReplaceAll(strChanged, ",", "")
	strChanged = strings.ReplaceAll(strChanged, "!", "")
	strChanged = strings.ReplaceAll(strChanged, "?", "")
	strChanged = strings.ReplaceAll(strChanged, "-", " ")
	strChanged = strings.ToLower(strChanged)
	strChanged = strings.TrimSpace(strChanged)
	strSlice := strings.Split(strChanged, " ")

	strMap := make(map[string]int)
	for _, val := range strSlice {
		if val != "" {
			strMap[val]++
		}
	}

	return strMap
}

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		println(err)
	}

	fmt.Println(countWords(input))
}
