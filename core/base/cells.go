package base

// A cell location is a list of coordinates corresponding to a location in the grid
type CellLocation []int

type Cells[T any] interface {
	RuleSet[T]
	GetNeighbors(cellLocation CellLocation) Neighborhood[T]
	Step() Cells[T]
}
