package main

type Config struct {
	simuDuration  int
	intersections int
	nStreets      int
	nCars         int
	bonusPoints   int
}

type Street struct {
	startIntersection int
	endIntersection   int
	name              string
	timeNeeded        int
}

type CarsPaths struct {
	nStreets    int
	streetNames []string
}

type Intersection struct {
	incomingStreets  map[string]Street
	outcomingStreets map[string]Street
}

func buildConfig(inputSet string) Config {
	parts := splitSpaces(inputSet)
	return Config{
		simuDuration:  toint(parts[0]),
		intersections: toint(parts[1]),
		nStreets:      toint(parts[2]),
		nCars:         toint(parts[3]),
		bonusPoints:   toint(parts[4]),
	}
}

func buildStreets(c Config, lines []string) ([]Street, map[string]Street, map[int]Intersection) {
	streets := make([]Street, c.nStreets)
	streetMap := make(map[string]Street, 0)

	intersectionMap := make(map[int]Intersection)

	for i := 0; i < c.nStreets; i++ {
		parts := splitSpaces(lines[i])
		street := Street{
			startIntersection: toint(parts[0]),
			endIntersection:   toint(parts[1]),
			name:              parts[2],
			timeNeeded:        toint(parts[3]),
		}

		intersectionA := intersectionMap[street.startIntersection]
		if intersectionA.outcomingStreets == nil {
			intersectionA.outcomingStreets = make(map[string]Street)
		}
		intersectionA.outcomingStreets[street.name] = street
		intersectionMap[street.startIntersection] = intersectionA

		intersectionB := intersectionMap[street.endIntersection]
		if intersectionB.incomingStreets == nil {
			intersectionB.incomingStreets = make(map[string]Street)
		}
		intersectionB.incomingStreets[street.name] = street
		intersectionMap[street.endIntersection] = intersectionB

		streets[i] = street
		streetMap[parts[2]] = street
	}
	return streets, streetMap, intersectionMap
}

func buildCarsPaths(c Config, lines []string) []CarsPaths {
	buildCars := make([]CarsPaths, c.nCars)

	for i := 0; i < c.nCars; i++ {
		parts := splitSpaces(lines[i])
		buildCars[i] = CarsPaths{
			nStreets:    toint(parts[0]),
			streetNames: parts[1:],
		}
	}

	return buildCars
}

func buildOutput(result int) string {
	return "42"
}
