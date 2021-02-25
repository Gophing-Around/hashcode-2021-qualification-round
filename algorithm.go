package main

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
	intersectionMap map[int]*Intersection) []output {

	out := make([]output, 0)

	for interseciontID, intersections := range intersectionMap {
		streetTimes := make([]StreetTime, 0)
		for _, street := range intersections.incomingStreets {
			streetTimes = append(streetTimes, StreetTime{
				name:               street.name,
				greenLigthDuration: config.simuDuration / len(intersections.incomingStreets),
			})
		}

		out = append(out, output{
			intersectionId: interseciontID,
			streetsTime:    streetTimes,
		})
	}

	return out
}
