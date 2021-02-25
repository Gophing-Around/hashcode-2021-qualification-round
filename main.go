package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fredmaggiowski/gowq"
)

func main() {
	files := []string{
		// Uncomment the line with the desired files (add other lines if needed)
		"a",
		// "a", "b", "c", "d", "e", "f",
		// "a", "b",
		// "a", "b", "e", "f",
		// "c",
		// "d",
	}

	gowq.NewWQ(10)

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in", fileName))

		input := buildInput(inputSet)
		printInputMetrics(input)

		result := algorithm(input)

		output := buildOutput(result)
		printResultMetrics(result)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
	}
}
