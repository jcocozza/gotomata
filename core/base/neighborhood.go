package base 

// returns a list of all of the neighborhoods for a given layer
type Neighborhoods[T any] func(layer int) [][]T
