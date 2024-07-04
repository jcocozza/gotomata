package core

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
