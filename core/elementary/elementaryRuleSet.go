package elementary

import "github.com/jcocozza/gotomata/core/base"

func btoi(b bool) int {
    if b { return 1} 
    return 0
}

type ElementaryRuleSet struct {
    ruleNumber uint8
}

func (ers ElementaryRuleSet) GetRule(neighborhood base.Neighborhood[bool]) base.Rule[bool] {
    left := btoi(neighborhood[0])
    center := btoi(neighborhood[1])
    right := btoi(neighborhood[2])
    
    rule := func(nb base.Neighborhood[bool]) bool {
        idx := left<<2 | center<<1 | right
        a := int(ers.ruleNumber >> idx & 1)
        return a == 1
    }
    return rule
}

