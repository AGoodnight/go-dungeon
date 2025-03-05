package creatures

import (
	"time"

	"github.com/google/uuid"
)

func CreateCreature() (c *Creature) {
	now := time.Now().UnixMicro() / int64(time.Millisecond)
	return &Creature{
		uuid.New().String(),
		false,
		10,
		1,
		0, 0,
		"Dark",
		now,
		now,
		"Monster",
	}
}
