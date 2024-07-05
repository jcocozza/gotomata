package core

import (
	"hash/fnv"
	"strconv"
)

// an n-dimensional coordinate on the grid
type Coordinate []int

/*
func (c Coordinate) hash() uint32 {
	hasher := fnv.New32a()
	for _, num := range c {
		bytes := []byte(fmt.Sprintf("%d", num))
		hasher.Write(bytes)
	}
	return hasher.Sum32()
}
*/

func (c Coordinate) hash() uint32 {
	hasher := fnv.New32a()
	for _, num := range c {
		hasher.Write([]byte(strconv.Itoa(num)))
	}
	return hasher.Sum32()
}

// a cell has a state of type T
type Cell[T comparable] struct {
	State      T
	Coordinate Coordinate
}

type BaseGrid[T comparable] struct {
	// a list of dimension sizes
	// i.e for a 4x4 grid, dimensions = []int{4,4}
	Dimensions []int
	Cells      *sparseCellGrid[T]
	// must specify default state to work with sparse grids.
	// the default state is the one that is not stored in the map
	DefaultState T
}

// if the state is the default state, delete the coordinate from the map
//
// otherwise add it to the map
func (bg *BaseGrid[T]) SetCell(state T, coordinate Coordinate) {
	if state == bg.DefaultState {
		bg.Cells.Delete(coordinate.hash())
	} else {
		bg.Cells.Set(coordinate.hash(), &Cell[T]{State: state, Coordinate: coordinate})
	}
}

// If the coordinate exists in the map, return it
//
// otherwise return a cell with the base state
func (bg *BaseGrid[T]) GetCell(coordinate Coordinate) *Cell[T] {
	if cell, exists := bg.Cells.Get(coordinate.hash()); exists {
		return cell
	}
	return &Cell[T]{State: bg.DefaultState, Coordinate: coordinate}
}

// get the cell by hash
//
// ONLY works for existing hashes
// Will panic if the hash is not in the map
func (bg *BaseGrid[T]) GetCellByHash(hash uint32) *Cell[T] {
	if cell, exists := bg.Cells.Get(hash); exists {
		return cell
	}
	panic("hash is not in map, something has gone very wrong")
}

func (bg *BaseGrid[T]) AllCoordinates() []Coordinate {
	totalCoords := 1
	for _, dim := range bg.Dimensions {
		totalCoords *= dim
	}
	coords := make([]Coordinate, totalCoords)

	var iter func(idx int, current Coordinate, dimension int) int
	iter = func(idx int, current Coordinate, dimension int) int {
		if dimension == len(bg.Dimensions) {
			coords[idx] = make(Coordinate, len(current))
			copy(coords[idx], current)
			return idx + 1
		}
		for i := 0; i < bg.Dimensions[dimension]; i++ {
			current[dimension] = i
			idx = iter(i, current, dimension+1)
		}
		return idx
	}

	iter(0, make(Coordinate, len(bg.Dimensions)), 0)
	return coords
}

// return all the neighbors of a coordinate
type GetNeighborsFunc func(coord Coordinate) []Coordinate

type Grid[T comparable] struct {
	*BaseGrid[T]
	GetNeighborCoordinates GetNeighborsFunc
}

func (g *Grid[T]) GetNeighbors(coord Coordinate) []*Cell[T] {
	neighbors := []*Cell[T]{}
	neighborCoords := g.GetNeighborCoordinates(coord)
	for _, co := range neighborCoords {
		cell := g.GetCell(co)
		neighbors = append(neighbors, cell)
	}
	return neighbors
}

// Return an empty grid with the same configuration as the one it comes from
func (g *Grid[T]) New() *Grid[T] {
	return &Grid[T]{
		BaseGrid: &BaseGrid[T]{
			Dimensions:   g.Dimensions,
			Cells:        NewSparseCellGrid[T](),
			DefaultState: g.DefaultState,
		},
		GetNeighborCoordinates: g.GetNeighborCoordinates,
	}
}
