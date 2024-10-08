package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2023/utils"
)

type SeedMap struct {
	source      string
	destination string
	values      SeedMapValues
}

type SeedMapValues struct {
	destinations []int
	sources      []int
	ranges       []int
}

func getSeeds(lines []string) []int {
	seeds := []int{}
	seedsLine := lines[0]

	seedsLine = strings.ReplaceAll(seedsLine, "seeds:", "")
	seedsLine = strings.TrimSpace(seedsLine)
	seedsStr := strings.Split(seedsLine, " ")

	for _, seedStr := range seedsStr {
		aux, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, aux)
	}

	return seeds
}

func getMaps(lines []string) []SeedMap {
	maps := []SeedMap{}

	mapsCoordinates := getMapsCoordinates(lines)
	for _, mapCoordinates := range mapsCoordinates {
		newMap := SeedMap{}
		newMapValues := SeedMapValues{}

		for i := mapCoordinates[0]; i <= mapCoordinates[1]; i++ {
			if i == mapCoordinates[0] {
				newMap.source = getMapSource(lines[i])
				newMap.destination = getMapDestination(lines[i])
			} else {
				values := strings.Split(lines[i], " ")

				newDestination, _ := strconv.Atoi(values[0])
				newSource, _ := strconv.Atoi(values[1])
				newRange, _ := strconv.Atoi(values[2])

				newMapValues.destinations = append(newMapValues.destinations, newDestination)
				newMapValues.sources = append(newMapValues.sources, newSource)
				newMapValues.ranges = append(newMapValues.ranges, newRange)
			}
		}
		newMap.values = newMapValues
		maps = append(maps, newMap)
	}

	return maps
}

func getMapSource(line string) string {
	destination := strings.Split(line, "-to-")
	strings.Trim(destination[0], " ")
	return destination[0]
}

func getMapDestination(line string) string {
	source := strings.Split(line, "-to-")
	source = strings.Split(source[1], " map:")
	trimmedSource := strings.Trim(source[0], " ")
	return trimmedSource
}

func getMapsCoordinates(lines []string) [][]int {
	emptyLinesIndex := []int{}
	mapsCoordinates := [][]int{}

	for i, line := range lines {
		if len(line) == 0 {
			emptyLinesIndex = append(emptyLinesIndex, i)
		}
	}

	for i := range emptyLinesIndex {
		if i+1 < len(emptyLinesIndex) {
			mapsCoordinates = append(
				mapsCoordinates,
				[]int{emptyLinesIndex[i] + 1, emptyLinesIndex[i+1] - 1})
		}
	}

	return mapsCoordinates
}

func partOne(seeds []int, seedMaps []SeedMap) int {
	// fmt.Printf("Seeds: %v\n", seeds)
	for i := range seeds {
		for _, seedMap := range seedMaps {
			fmt.Printf("%s-to-%s map:\n", seedMap.source, seedMap.destination)

			// Loop through each line
			for j := 0; j < len(seedMap.values.destinations); j++ {
				fmt.Printf("%d %d %d\n\n",
					seedMap.values.destinations[j],
					seedMap.values.sources[j],
					seedMap.values.ranges[j])

				// Loop through each range value
				for k := 0; k < seedMap.values.ranges[j]; k++ {
					fmt.Printf("%d -> %d\n", seedMap.values.sources[j]+k, seedMap.values.destinations[j]+k)
					if seeds[i] == seedMap.values.sources[j]+k {
						seeds[i] = seedMap.values.destinations[j] + k
					}
				}
			}
			break
		}
		break
	}

	return 0
}

func main() {
	lines := utils.GetLinesFile("example.txt")
	seeds := getSeeds(lines)
	maps := getMaps(lines)

	resultPartOne := partOne(seeds, maps)
	fmt.Printf("Part 1: %d\n", resultPartOne)

	// resultPartTwo := partTwo(seeds, maps)
	// fmt.Printf("Part 2: %d\n", resultPartTwo)
}
