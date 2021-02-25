package main

type Config struct {
	simuDuration  int
	intersections int
	nStreets      int
	nCars         int
	bonusPoints   int
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

func buildStreets(c Config, lines []string) []Street {
	streets := make([]Street, c.nStreets)

	for i := 0; i < c.nStreets; i++ {
		parts := splitSpaces(lines[i])
		streets[i] = Street{
			startIntersection: toint(parts[0]),
			endIntersection:   toint(parts[1]),
			name:              parts[2],
			timeNeeded:        toint(parts[3]),
		}
	}
	return streets
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
