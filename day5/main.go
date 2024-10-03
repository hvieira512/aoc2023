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
	sources      []int
	destinations []int
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

	// loop while I dont find a blank new line
	for i, line := range lines {
		// no need to read the seed
		if i < 2 {
			continue
		}

		// new map detected
		if strings.Contains(line, "map") {
			// get source and destination
			newMapSource := getMapSource(line)
			newMapDestination := getMapDestination(line)
			fmt.Println(newMapSource, newMapDestination)

			// get values
		}
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

func main() {
	lines := utils.GetLinesFile("example.txt")
	seeds := getSeeds(lines)
	maps := getMaps(lines)

	fmt.Printf("Seeds: %v\n", seeds)
	fmt.Printf("Maps: %v\n", maps)
}
