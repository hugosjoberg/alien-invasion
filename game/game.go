package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hugosjoberg/alien-invasion/cmap"
)

// Game struct holds information used to run the games, the map of the cities,
// and the number of aliens left and the number of iterations of the game.
type Game struct {
	cityMap map[string]cmap.CityNode
	moves   int
	aliens  int
}

// placeAliensCities places the aliens in random cities on the map.
func (g Game) placeAliensCities() {
	rand.Seed(time.Now().UnixNano())
	cities := []string{}
	for cityName := range g.cityMap {
		cities = append(cities, cityName)
	}
	for i := 0; i < g.aliens; i++ {
		randCity := cities[rand.Intn(len(cities))]
		g.cityMap[randCity].Aliens[fmt.Sprintf("Alien %d", i)] = cmap.Alien{Moves: 0}
	}
}

// destroyCity destroys a city and all references to the city.
func (g Game) destroyCity(city string) {
	references := g.cityMap[city]
	if cn, ok := g.cityMap[references.North]; ok {
		cn.South = ""
		g.cityMap[references.North] = cn
	}
	if cn, ok := g.cityMap[references.South]; ok {
		cn.North = ""
		g.cityMap[references.South] = cn
	}
	if cn, ok := g.cityMap[references.East]; ok {
		cn.West = ""
		g.cityMap[references.East] = cn
	}
	if cn, ok := g.cityMap[references.West]; ok {
		cn.East = ""
		g.cityMap[references.West] = cn
	}
	delete(g.cityMap, city)
}

// moveAliens moves each alien to a neighbor city
func (g *Game) moveAliens() {
	rand.Seed(time.Now().UnixNano())
	for _, city := range g.cityMap {
		cities := []string{}
		if city.North != "" {
			cities = append(cities, city.North)
		}
		if city.South != "" {
			cities = append(cities, city.South)
		}
		if city.East != "" {
			cities = append(cities, city.East)
		}
		if city.West != "" {
			cities = append(cities, city.West)
		}
		for alienName, alien := range city.Aliens {
			// check if alien can make move this turn
			if alien.Moves == g.moves {
				// check if the alien is trapped
				if len(cities) <= 0 {
					// the alien is trapped, let's kill it
					delete(city.Aliens, alienName)
					g.aliens -= 1
				} else {
					alien.Moves += 1
					// move the alien to a random city
					randCity := cities[rand.Intn(len(cities))]
					g.cityMap[randCity].Aliens[alienName] = alien
					delete(city.Aliens, alienName)
				}
			}
		}
	}
	g.moves += 1
}

// alienFight checks all the cities for more than 1 alien, and make the aliens fight if there are more than 1 alien.
func (g *Game) alienFight() {
	for cityName, city := range g.cityMap {
		// check if one or more aliens occupy a city
		if len(city.Aliens) > 1 {
			g.aliens -= len(city.Aliens)
			aliens := []string{}
			for alienName := range city.Aliens {
				aliens = append(aliens, alienName)
				delete(city.Aliens, alienName)
			}
			fmt.Printf("Aliens %s, threw an epic fight in %s, all the aliens died and %s got leveled\n", strings.Join(aliens, ", "), cityName, cityName)
			g.destroyCity(cityName)
		}
	}
}

// Run runs the game until there are no more cities for the aliens
// to destroy or if the aliens are not able to move to a new city.
func Run(mapPath string, aliens int) error {
	cityMap, err := cmap.New(mapPath)
	if err != nil {
		return err
	}
	game := Game{cityMap: cityMap, moves: 0, aliens: aliens}
	game.placeAliensCities()
	for {
		// Let the aliens fight
		game.alienFight()
		// move the aliens
		game.moveAliens()
		if game.moves >= 10000 || game.aliens <= 1 {
			break
		}
	}
	fmt.Printf("Game ended after %d moves\n", game.moves)
	fmt.Printf("Cities left: %d\n", len(game.cityMap))
	fmt.Printf("Aliens left: %d\n", game.aliens)
	return nil
}
