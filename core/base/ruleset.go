package base

// A rule takes a neighborhood to a state
type Rule[T any] func(neighborhood Neighborhood[T]) T 

// A rule set is a list of rules governing the next state based on a neighborhood
type RuleSet[T any] interface {
    GetRule(neighborhood Neighborhood[T]) Rule[T]
}
