package creatures

import (
	"fmt"
	"slices"
	"testing"
)

func TestAttack(t *testing.T) {
	tally := map[string][2]int{}
	critters := CreateCreatures(5)
	results := []bool{}
	for _, creature := range critters {
		was := creature.HitPoints
		Attack(&creature, 1)
		is := creature.HitPoints
		tally[creature.Name] = [2]int{was, is}
		fmt.Println("is", is, "was:", was, creature.IsAlive(), tally[creature.Name])
	}
	for _, member := range tally {
		if member[0] == member[0]-1 {
			results = append(results, true)
		}
	}
	if slices.Contains(results, false) {
		t.Errorf("Hit Points not affected by Attack")
	}
}
