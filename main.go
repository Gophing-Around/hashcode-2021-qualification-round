package main

import (
	"fmt"

	"github.com/fredmaggiowski/gowq"
)

func main() {
	files := []string{
		// Uncomment the line with the desired files (add other lines if needed)
		"a",
		// "a",
		// "a",
		// "a",
		// "c",
		// "d",
	}

	gowq.NewWQ(10)

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.txt", fileName))

		fileLines := splitNewLines(inputSet)

		config := buildConfig(fileLines[0])
		// fmt.Printf("%+v\n", config)

		streets := buildStreets(config, fileLines[1:])
		// fmt.Printf("%+v\n", streets)
		carsPaths := buildCarsPaths(config, fileLines[1+config.nStreets:])
		// fmt.Printf("%+v\n", carsPaths)

		// printInputMetrics(input)
		// result := algorithm(input)

		// output := buildOutput(result)
		// printResultMetrics(result)

		// ioutil.WriteFile(fmt.Sprintf("./result/%s.txt", fileName), []byte(output), 0644)
	}
}
