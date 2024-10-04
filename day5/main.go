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
			mapsCoordinates = append(mapsCoordinates, []int{emptyLinesIndex[i] + 1, emptyLinesIndex[i+1] - 1})
		}
	}

	return mapsCoordinates
}

func renderSeedMaps(maps []SeedMap) {
	for _, seedMap := range maps {
		fmt.Printf("%s-to-%s\n", seedMap.source, seedMap.destination)
		for i := 0; i < len(seedMap.values.ranges); i++ {
			fmt.Printf(
				"%d %d %d\n",
				seedMap.values.destinations[i],
				seedMap.values.sources[i],
				seedMap.values.ranges[i],
			)
		}
		fmt.Println()
	}
}

func main() {
	lines := utils.GetLinesFile("example.txt")
	seeds := getSeeds(lines)
	maps := getMaps(lines)

	fmt.Printf("Seeds: %v\n", seeds)
	renderSeedMaps(maps)
}
