package core

//import "fmt"

type CellularAutomata[T comparable] struct {
	Grid    *Grid[T]
	RuleSet RuleSet[T]
	Steps   int
}

func (ca *CellularAutomata[T]) Step() {
	cellsToCheck := make(CellSet[T])
	for _, key := range ca.Grid.Cells.GetAllKeys() {
		// add the cell to the set of cells to check
		mainCell := ca.Grid.BaseGrid.GetCellByHash(key)
		cellsToCheck.Add(key, mainCell)
		// but we also need to check all neighbors of the cell too
		neighbors := ca.Grid.GetNeighbors(mainCell.Coordinate)
		for _, cell := range neighbors {
			cellsToCheck.Add(cell.Coordinate.hash(), cell)
		}
	}

	newGrid := ca.Grid.New()

	for _, cell := range cellsToCheck {
		neighbors := ca.Grid.GetNeighbors(cell.Coordinate)
		next := ca.RuleSet(cell, neighbors)
		newGrid.SetCell(next.State, next.Coordinate)
	}

	ca.Grid = newGrid
}

func (ca *CellularAutomata[T]) Stepp() {
	newGrid := ca.Grid.New()

	var processer = func(shard int, cells map[uint64]*Cell[T]) {
		localCellsToCheck := make(CellSet[T])
		for key, cell := range cells {
			localCellsToCheck.Add(key, cell)
			neighbors := ca.Grid.GetNeighbors(cell.Coordinate)
			for _, neighbor := range neighbors {
				localCellsToCheck.Add(neighbor.Coordinate.hash(), neighbor)
			}
		}
		for _, cell := range localCellsToCheck {
			neighbors := ca.Grid.GetNeighbors(cell.Coordinate)
			next := ca.RuleSet(cell, neighbors)
//			fmt.Printf("%v | NBHD: %v\n", cell, neighbors)
			newGrid.SetCell(next.State, next.Coordinate)
		}
	}
	ca.Grid.Cells.ProcessShard(processer)
	ca.Grid = newGrid
}

// apply the ruleset only at the passed coordinate
//
// returns the coordinate of where the next cell goes
//
// this is used to implement random walks (since we only keep track of the most recent coordinate)
func (ca *CellularAutomata[T]) StepHead(coordinate Coordinate) Coordinate {
	cell := ca.Grid.GetCell(coordinate)
	neighbors := ca.Grid.GetNeighbors(coordinate)
	next := ca.RuleSet(cell, neighbors)
	ca.Grid.SetCell(next.State, next.Coordinate)
	return next.Coordinate
}
