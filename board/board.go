package board

import "dungeon/creatures"

func Attack(c creatures.CreatureLifeCycle, taken int) (int, bool) {
	return c.Damage(taken)
}
