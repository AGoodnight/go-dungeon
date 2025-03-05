package creatures

type CreatureLifeCycle interface {
	Summon()
	Kill()
	IsAlive() bool
	Damage(taken int) (int, bool)
}

type Creature struct {
	ID           string `json:"id"`
	Alive        bool   `json:"alive"`
	HitPoints    int    `json:"hitPoints"`
	Initiative   int    `json:"initiative"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Category     string `json:"category"`
	TimeCreated  int64  `json:"timeCreated"`
	TimeModified int64  `json:"timeModified"`
	Name         string `json:"name"`
}

type Player struct {
	Creature
	Success bool `json:"success"`
}

func (c *Creature) Summon() {
	c.Alive = true
}

func (c *Creature) Kill() {
	c.Alive = false
}

func (c *Creature) IsAlive() bool {
	return c.Alive
}

func (c *Creature) Damage(taken int) (int, bool) {
	c.HitPoints -= taken
	if c.HitPoints < 1 {
		c.Alive = false
	}
	return c.HitPoints, c.Alive
}
