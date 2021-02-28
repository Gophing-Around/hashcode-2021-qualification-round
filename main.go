package main

import (
	"fmt"
	"io/ioutil"
	"sort"

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

		streets, streetsMap, intersectionMap, intersectionsList := buildStreets(config, fileLines[1:])
		carsPaths := buildCarsPaths(config, fileLines[1+config.nStreets:])

		intersectionsList = sortIntersections(intersectionsList)

		// printInputMetrics(input)
		outputs := algorithm(config, streets, carsPaths, streetsMap, intersectionMap, intersectionsList)

		result := buildOutput(outputs)
		// printResultMetrics(result)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.txt", fileName), []byte(result), 0644)
	}
}

func sortIntersections(list []*Intersection) []*Intersection {
	sort.Slice(list, func(i, j int) bool {
		return list[i].arrivingCars > list[j].arrivingCars
	})
	return list
}
