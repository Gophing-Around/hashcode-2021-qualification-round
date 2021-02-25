package main

func algorithm2(
	config Config,
	streets []*Street,
	carsPaths []*CarsPaths,
	streetsMap map[string]*Street,
	intersectionMap map[int]*Intersection,
	intersectionsList []*Intersection,
) []output {
	out := make([]output, 0)

	for tick := 0; tick < config.simuDuration; {

		tick++
	}

	for _, intersection := range intersectionsList {
		streetTimes := make([]StreetTime, 0)
		totScore := 0

		for _, street := range intersection.incomingStreets {
			totScore += int(street.score)
			// fmt.Printf("STREET %s - score: %d.\n", street.name, street.score)
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
