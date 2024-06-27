package base

// A grid is a slice of slices
type Grid[T any] struct {
	Dimensions int
	Sizes      []int
	Data       []T
}

// Create a new grid
//
// sizes are a list of ints specifing the number of cells in a given direction
//
// e.g. size = []int{50,50} will be a 50 x 50 grid
func NewGrid[T any](sizes []int) *Grid[T] {
	dimensions := len(sizes)
		
	totalCells := 1 
	for _, size := range sizes {
		totalCells *= size
	}
	data := make([]T, totalCells)

	return &Grid[T]{
		Dimensions: dimensions,
		Sizes:      sizes,
		Data:       data,
	}

}

// Compute the index for the cell based on a cell location
func (g *Grid[T]) Index(cellLocation CellLocation) int {
	index := 0
	stride := 1
	for i := g.Dimensions - 1; i >= 0; i-- {
		index += cellLocation[i] * stride
		stride *= g.Sizes[i]
	}
	return index
}

// Coords calculates the coordinates in the n-dimensional grid for a given linear index
func (g *Grid[T]) Coords(index int) []int {
	coords := make([]int, g.Dimensions)

	for i := g.Dimensions - 1; i >= 0; i-- {
		coords[i] = index % g.Sizes[i]
		index /= g.Sizes[i]
	}
	

	return coords
}

func (g *Grid[T]) GetValue(idx int) T {
	return g.Data[idx]
}

func (g *Grid[T]) SetValue(idx int, value T) {
	g.Data[idx] = value
}

func (g *Grid[T]) SetValueByCoord(value T, coord CellLocation) {
	idx := g.Index(coord)
	g.Data[idx] = value
}
