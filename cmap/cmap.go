// Package cmap implements parsing of a map file.
package cmap

import (
	"bufio"
	"os"
	"strings"
)

// CityNode struct holds the information about a city
type CityNode struct {
	North  string
	South  string
	East   string
	West   string
	Aliens map[string]Alien
}

// Alien struct keeps the information about an Alien
type Alien struct {
	Moves int
}

// parseLine parses a line and returns the city name and the city nodes.
func parseLine(line []string) (string, CityNode) {
	cityName := line[0]
	node := CityNode{Aliens: make(map[string]Alien)}
	for _, s := range line {
		switch {
		case strings.Contains(s, "north="):
			node.North = strings.Split(s, "=")[1]
		case strings.Contains(s, "south="):
			node.South = strings.Split(s, "=")[1]
		case strings.Contains(s, "east="):
			node.East = strings.Split(s, "=")[1]
		case strings.Contains(s, "west="):
			node.West = strings.Split(s, "=")[1]
		}
	}
	return cityName, node
}

// New returns the game map.
func New(path string) (map[string]CityNode, error) {
	cityMap := make(map[string]CityNode)
	file, err := os.Open(path)
	if err != nil {
		return cityMap, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) > 0 {
			cityName, cityNode := parseLine(line)
			if _, ok := cityMap[cityNode.North]; cityNode.North != "" && !ok {
				cityMap[cityNode.North] = CityNode{South: cityName, Aliens: make(map[string]Alien)}
			}
			if _, ok := cityMap[cityNode.South]; cityNode.South != "" && !ok {
				cityMap[cityNode.South] = CityNode{North: cityName, Aliens: make(map[string]Alien)}
			}
			if _, ok := cityMap[cityNode.East]; cityNode.East != "" && !ok {
				cityMap[cityNode.East] = CityNode{West: cityName, Aliens: make(map[string]Alien)}
			}
			if _, ok := cityMap[cityNode.West]; cityNode.West != "" && !ok {
				cityMap[cityNode.West] = CityNode{East: cityName, Aliens: make(map[string]Alien)}
			}
			cityMap[cityName] = cityNode
		}
	}

	if err := scanner.Err(); err != nil {
		return make(map[string]CityNode), err
	}
	return cityMap, nil
}
