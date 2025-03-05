package main

import (
	"dungeon/board"
	"dungeon/creatures"
	"fmt"
)

func main() {
	creature := creatures.CreateCreature()
	creature.Summon()
	for creature.HitPoints > 0 {
		board.Attack(creature, 1)
		fmt.Println(creature.IsAlive())
	}
}
