package game

import (
	"testing"

	"github.com/hugosjoberg/alien-invasion/cmap"
)

func TestDestroyCity(t *testing.T) {
	cityMap := make(map[string]cmap.CityNode)
	cityMap["Irvine"] = cmap.CityNode{North: "Los-Angeles", East: "Las Vegas"}
	cityMap["Los-Angeles"] = cmap.CityNode{South: "Irvine"}
	cityMap["Las Vegas"] = cmap.CityNode{West: "Irvine"}
	game := Game{cityMap: cityMap}
	game.destroyCity("Irvine")
	if _, ok := game.cityMap["Irvine"]; ok {
		t.Errorf("expect city %s to be destroyed", "Irvine")
	}
	if game.cityMap["Los-Angeles"].South != "" {
		t.Errorf("expect reference to city %s to be destroyed", "Irvine")
	}
	if game.cityMap["Las Vegas"].South != "" {
		t.Errorf("expect reference to city %s to be destroyed", "Irvine")
	}
}

func TestMoveAliens(t *testing.T) {
	cityMap := make(map[string]cmap.CityNode)
	aliens := make(map[string]cmap.Alien)
	aliens["alien1"] = cmap.Alien{Moves: 0}
	cityMap["Irvine"] = cmap.CityNode{North: "Los-Angeles", Aliens: aliens}
	cityMap["Los-Angeles"] = cmap.CityNode{South: "Irvine", Aliens: make(map[string]cmap.Alien)}
	game := Game{cityMap: cityMap, moves: 0}
	game.moveAliens()
	if len(game.cityMap["Irvine"].Aliens) > 0 {
		t.Errorf("expect no aliens in Irvine, found %d", len(game.cityMap["Irvine"].Aliens))
	}
	if len(game.cityMap["Los-Angeles"].Aliens) == 0 {
		t.Errorf("expect at least one alien in Los Angeles, found %d", len(game.cityMap["Los-Angeles"].Aliens))
	}
	game.moveAliens()
	if len(game.cityMap["Los-Angeles"].Aliens) > 0 {
		t.Errorf("expect no aliens in Los-Angeles, found %d", len(game.cityMap["Los-Angeles"].Aliens))
	}
	if len(game.cityMap["Irvine"].Aliens) == 0 {
		t.Errorf("expect at least one alien in Irvine, found %d", len(game.cityMap["Irvine"].Aliens))
	}
}

func TestAlienFight(t *testing.T) {
	cityMap := make(map[string]cmap.CityNode)
	aliens := make(map[string]cmap.Alien)
	aliens["alien1"] = cmap.Alien{Moves: 0}
	cityMap["Irvine"] = cmap.CityNode{Aliens: aliens}
	game := Game{cityMap: cityMap, moves: 0}

	// only 1 alien in Irvine, no fighting to be done
	game.alienFight()
	if len(game.cityMap["Irvine"].Aliens) > 0 {
		t.Errorf("expect 1 alien, found %d", len(game.cityMap["Irvine"].Aliens))
	}

	// 2 aliens in Irvine, expect them to fight and die, in their death they should also bring down Irvine
	aliens["alien2"] = cmap.Alien{Moves: 0}
	game.alienFight()
	if len(game.cityMap["Irvine"].Aliens) != 0 {
		t.Errorf("expect 0 alien, found %d", len(game.cityMap["Irvine"].Aliens))
	}
	if _, ok := game.cityMap["Irvine"]; ok {
		t.Errorf("expect city %s to be destroyed", "Irvine")
	}
}
