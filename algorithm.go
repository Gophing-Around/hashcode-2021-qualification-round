package main

import "sort"

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
		var lastStreetName string
		previouslyEnqueued := -1
		for _, street := range carPath.streetNames {
			street := streetsMap[street]
			totTime += street.timeNeeded

			if previouslyEnqueued < 0 {
				previouslyEnqueued = street.firstQueue
				street.firstQueue++
			}

			street.passingCars++

			intersectionId = street.endIntersection
			lastStreetName = street.name
			streetsMap[street.name] = street
		}

		if totTime + previouslyEnqueued < 9*config.simuDuration/10 {
			street := streetsMap[lastStreetName]
			intersection := intersectionMap[intersectionId]

			street.arrivingCars++
			street.passingCars--
			intersection.arrivingCars++

			streetsMap[lastStreetName] = street
			intersectionMap[intersectionId] = intersection
		} else {
			for _, street := range carPath.streetNames {
				street := streetsMap[street]
				street.passingCars--
				streetsMap[street.name] = street
			}
		}
	}

	visited := make(map[int]bool)
	for _, intersection := range intersectionsList {
		dfs(
			visited,
			config.simuDuration,
			intersection,
			0,
			intersection,
			intersectionMap,
		)
	}

	for _, intersection := range intersectionsList {
		streetTimes := make([]StreetTime, 0)
		totScore := 0
		totPassingCars := 0

		for _, street := range intersection.incomingStreets {
			totScore += int(street.score)
			totPassingCars += street.passingCars
		}
		if totScore == 0 || totPassingCars == 0 {
			continue
		}

		incomingStreets := make([]*Street, 0)
		for _, street := range intersection.incomingStreets {
			incomingStreets = append(incomingStreets, street)
		}

		sort.Slice(incomingStreets, func(i, j int) bool {
			valueA := incomingStreets[i]
			valueB := incomingStreets[j]
			return valueA.passingCars > valueB.passingCars
		})

		for _, street := range incomingStreets {
			if street.score == 0 || street.passingCars == 0 {
				continue
			}

			a := street.passingCars * config.simuDuration / ( totPassingCars * 100 )
			if a <= 0 || a > config.simuDuration {
				a = 1
			}

			streetTimes = append(streetTimes, StreetTime{
				name:               street.name,
				greenLigthDuration: a, // / len(intersection.incomingStreets),
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
) {
	if score <= intersection.maxScore {
		return
	}
	intersection.maxScore = score

	if visited := visited[intersection.id]; visited {
		return
	}

	visited[intersection.id] = true

	// score += intersection.arrivingCars // TODO

	for streetName, incomingStreet := range intersection.incomingStreets {
		if incomingStreet.passingCars == 0 {
			continue
		}

		streetScore := score + 3*incomingStreet.arrivingCars + incomingStreet.passingCars - incomingStreet.timeNeeded
		dfs(
			visited,
			remainingTime-incomingStreet.timeNeeded,
			// incomingStreet,
			intersectionMap[incomingStreet.startIntersection],
			streetScore,

			startIntersection,
			intersectionMap,
		)
		
		incomingStreet.score = streetScore
		intersection.incomingStreets[streetName] = incomingStreet
	}

	visited[intersection.id] = false
}

func pickBestIntersection(intersectionsList []*Intersection) *Intersection {
	return intersectionsList[0]
}

type IntersectionNode struct {
	data *Intersection

	nextIntersections []*Intersection
}
