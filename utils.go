package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(source string) string {
	in, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}
	return string(in)
}

func toint(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return res
}

func splitNewLines(rawContent string) []string {
	return strings.Split(rawContent, "\n")
}

func splitSpaces(lineContent string) []string {
	return strings.Split(lineContent, " ")
}

func joinWithSpaces(partials []string) string {
	return strings.Join(partials, " ")
}

/* useful pieces of code:

***** [1] sorting *****
sort.Slice(arrayOfValues, func(i, j int) bool {
	valueA := arrayOfValues[i]
	valueB := arrayOfValues[j]
	return valueA.score > valueB.score
})

***** [2] Format output string with Sprintf *****
output := ""
output += fmt.Sprintf("%s %d\n", stringValue, intValue)

*/
