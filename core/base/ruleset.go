package base 

// a rule determines what a neighborhood maps to
type Rule[T any] func(neighborhood []T) T

// Every list of rules must have a way of producing which rule to use based on the passed neighborhood
type RuleSet[T any] interface {
    GetRule(neighborhood []T) Rule[T]
}

