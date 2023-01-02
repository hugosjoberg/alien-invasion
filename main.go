package main

import (
	"flag"
	"log"

	"github.com/hugosjoberg/alien-invasion/game"
)

func main() {
	aliens := flag.Int("n", 2, "number of aliens")
	mapPath := flag.String("p", "example/map2.txt", "path to the map")
	flag.Parse()
	err := game.Run(*mapPath, *aliens)
	if err != nil {
		log.Fatal(err)
	}
}
