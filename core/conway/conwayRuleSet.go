package conway

import "github.com/jcocozza/gotomata/core/base"

/*

A neighborhood:
0 1 2
3 C 5
6 7 8
-------
=> [0,1,2,3,C,5,6,7,8]
*/

var (
	nothing base.Rule[bool] = func(neighborhood base.Neighborhood[bool]) bool {
		return false
	}
	// less then 2 live neighbors -> dies (underpopulation)
	underpopulation base.Rule[bool] = func(neighborhood base.Neighborhood[bool]) bool {
		return false
	}
	// 2 or 3 live neighborss -> lives (generation persists)
	generational base.Rule[bool] = func(neighborhood base.Neighborhood[bool]) bool {
		return true
	}
	// 3+  live neighbors -> dies (overpopulation)
	overpopulation base.Rule[bool] = func(neighborhood base.Neighborhood[bool]) bool {
		return false 
	}
	// dead & 3 live neighbors -> live (reproduction)
	reproduction base.Rule[bool] = func(neighborhood base.Neighborhood[bool]) bool {
		return true
	}
)

// implements base.RuleSet
type ConwayRuleSet struct{}

func (crs ConwayRuleSet) GetRule(neighborhood []bool) base.Rule[bool] {
	centerAlive := neighborhood[4]

	totalLive := 0
	for i, nb := range neighborhood {
		if nb && i != 4 {
			totalLive += 1
		}
	}

	switch {
	case totalLive < 2 && centerAlive:
		return underpopulation
	case (totalLive == 2 || totalLive == 3) && centerAlive:
		return generational
	case totalLive > 3 && centerAlive:
		return overpopulation
	case !centerAlive && totalLive == 3:
		return reproduction
	default:
		return nothing
	}
}
