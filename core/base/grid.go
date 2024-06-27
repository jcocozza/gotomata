package base

// A grid is a slice of slices
type Grid[T any] struct {
	Dimensions int
	Sizes      []int
	Data       []T
}

func NewGrid[T any](sizes []int) *Grid[T] {
	dimensions := len(sizes)

	totalCells := 0
	for _, size := range sizes {
		totalCells += size
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
	for i := len(g.Sizes) - 1; i >= 0; i-- {
		index += cellLocation[i] * stride
		stride *= g.Sizes[i]
	}
	return index
}

func (g *Grid[T]) GetValue(idx int) T {
	return g.Data[idx]
}

func (g *Grid[T]) SetValue(idx int, value T) {
	g.Data[idx] = value 
}
