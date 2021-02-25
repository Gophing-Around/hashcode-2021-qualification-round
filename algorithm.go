package main

import (
	"fmt"
	"sort"
)

type StreetTime struct {
	name               string
	greenLigthDuration int
}

type output struct {
	intersectionId int
	intersection   Intersection
	streetsTime    []StreetTime
}

func algorithm(
	config Config,
	streets []*Street,
	carsPaths []*CarsPaths,
	streetsMap map[string]*Street,
	intersectionMap map[int]*Intersection,
	intersectionsList []*Intersection,
) []output {
	out := make([]output, 0)

	for _, carPath := range carsPaths {
		var intersectionId int
		var totTime int
		for _, street := range carPath.streetNames {
			street := streetsMap[street]
			totTime += street.timeNeeded

			intersectionId = street.endIntersection
		}

		if totTime <= config.simuDuration {
			intersection := intersectionMap[intersectionId]
			intersection.arrivingCars++
			intersectionMap[intersectionId] = intersection
		}
	}

	sort.Slice(intersectionsList, func(i, j int) bool {
		return intersectionsList[i].arrivingCars > intersectionsList[j].arrivingCars
	})

	visited := make(map[int]bool)
	for _, intersection := range intersectionsList {
		dfs(
			visited,
			config.simuDuration,
			intersection,
			intersection.arrivingCars,
			intersection,
			intersectionMap,
		)
	}

	for _, intersection := range intersectionsList {
		streetTimes := make([]StreetTime, 0)
		totScore := 0

		for _, street := range intersection.incomingStreets {
			totScore += int(street.score)
			fmt.Printf("STREET %s - score: %d.\n", street.name, street.score)
		}
		if totScore == 0 {
			continue
		}

		for _, street := range intersection.incomingStreets {
			if street.score == 0 {
				continue
			}

			streetTimes = append(streetTimes, StreetTime{
				name:               street.name,
				greenLigthDuration: int(street.score) * config.simuDuration / totScore, // / len(intersection.incomingStreets),
			})
		}

		out = append(out, output{
			intersectionId: intersection.id,
			streetsTime:    streetTimes,
		})
	}

	return out
}

func dfs(
	visited map[int]bool,
	remainingTime int,
	// incomingStreet *Street,
	intersection *Intersection,
	score int,
	startIntersection *Intersection,

	intersectionMap map[int]*Intersection,
) int {
	if visited := visited[intersection.id]; visited || remainingTime < 0 {
		return score
	}

	visited[intersection.id] = true

	// score += intersection.arrivingCars // TODO

	for _, streetName := range intersection.incomingStreetsNames {

		incomingStreet := intersection.incomingStreets[streetName]

		streetScore := dfs(
			visited,
			remainingTime-incomingStreet.timeNeeded,
			// incomingStreet,
			intersectionMap[incomingStreet.startIntersection],
			score+int(intersectionMap[incomingStreet.startIntersection].arrivingCars),

			startIntersection,
			intersectionMap,
		)
		// score += streetScore
		incomingStreet.score = streetScore
		intersection.incomingStreets[streetName] = incomingStreet
	}

	// visited[intersection.id] = false
	return score
}

func pickBestIntersection(intersectionsList []*Intersection) *Intersection {
	return intersectionsList[0]
}

type IntersectionNode struct {
	data *Intersection

	nextIntersections []*Intersection
}
