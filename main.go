package main

import (
	"dungeon/board"
	"dungeon/creatures"
	"fmt"
)

func main() {
	critters := creatures.CreateCreatures(4)
	for _, creature := range critters {
		was := creature.HitPoints
		board.Attack(&creature, 1)
		fmt.Println("is", creature.HitPoints, "was:", was, creature.IsAlive())
	}
}
