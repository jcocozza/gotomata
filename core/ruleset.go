package core

/*
type Rule[T any] func(cell *Cell[T], neighbors []*Cell[T]) *Cell[T]

// A Ruleset is a set of rules and a way of knowing which one to apply
//
// (Technically it could be done in 1 function, but this should keep things easier for complex rules
type RuleSet[T any] struct{
        // A list of rules that can be applied
            Rules []Rule[T]
                // Which rule to use based on the conditions
                    GetRule func(cell *Cell[T], neighbors []*Cell[T]) Rule[T]
                }
*/

type RuleSet[T comparable] func(cell *Cell[T], neighbors []*Cell[T]) *Cell[T]
