package creatures

func Attack(c CreatureLifeCycle, taken int) (int, bool) {
	return c.Damage(taken)
}
