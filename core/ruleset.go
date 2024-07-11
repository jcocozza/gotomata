package core

// A Ruleset is a set of rules and a way of knowing which one to apply
//
// Technically, this should be a list of functions, but often it is much simpler to encapsulate checking criteria and output into a single function
type RuleSet[T comparable] func(cell *Cell[T], neighbors []*Cell[T]) *Cell[T]
