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
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	gowq.NewWQ(10)

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.txt", fileName))

		fileLines := splitNewLines(inputSet)

		config := buildConfig(fileLines[0])
		// fmt.Printf("%+v\n", config)

		streets, streetsMap, intersectionMap := buildStreets(config, fileLines[1:])
		// fmt.Printf("%+v\n", intersectionMap)
		// fmt.Printf("%+v\n", streets)
		carsPaths := buildCarsPaths(config, fileLines[1+config.nStreets:])
		// fmt.Printf("%+v\n", carsPaths)

		// printInputMetrics(input)
		outputs := algorithm(config, streets, carsPaths, streetsMap, intersectionMap)

		result := fmt.Sprintf("%d\n", len(outputs))
		for _, out := range outputs {
			result += fmt.Sprintf("%d\n", out.intersectionId)
			result += fmt.Sprintf("%d\n", len(out.streetsTime))
			for _, streetTime := range out.streetsTime {
				result += fmt.Sprintf("%s %d\n", streetTime.name, streetTime.greenLigthDuration)

			}
		}

		// output := buildOutput(result)
		// printResultMetrics(result)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.txt", fileName), []byte(result), 0644)
	}
}
