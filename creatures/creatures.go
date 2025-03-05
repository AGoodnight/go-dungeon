package creatures

import (
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

func CreateCreatures(amt int) (c []Creature) {
	now := time.Now().UnixMicro() / int64(time.Millisecond)
	creatures := []Creature{}

	for amt > 0 {
		next := &Creature{
			ID:           uuid.New().String(),
			Alive:        true,
			HitPoints:    rand.IntN(10),
			Initiative:   rand.IntN(3),
			X:            rand.IntN(20),
			Y:            rand.IntN(20),
			Category:     "Baddie",
			TimeCreated:  now,
			TimeModified: now,
			Name:         "Monster",
		}
		creatures = append(creatures, *next)
		amt--
	}

	return creatures
}
