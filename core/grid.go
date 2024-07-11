package core

import (
	"fmt"
	//"hash/fnv"
	//"strconv"
)

// an n-dimensional coordinate on the grid
type Coordinate []int

/*
func (c Coordinate) hash() uint64 {
	var hash uint64 = 14695981039346656037 // FNV offset basis
	for _, num := range c {
		hash ^= uint64(num)
		hash *= 1099511628211 // FNV prime
	}
	return hash
}
*/

func (c Coordinate) hash() uint64 {
	var hash uint64 = 14695981039346656037 // FNV offset basis
	for i, num := range c {
		// Incorporate the index to differentiate [1,2,3] from [3,2,1]
		hash ^= uint64(i)
		hash *= 1099511628211 // FNV prime
		hash ^= uint64(num)
		hash *= 1099511628211 // FNV prime
	}
	return hash
}

func (c Coordinate) String() string {
	s := "["
	for i, v := range c {
		if i == len(c)-1 {
			s += fmt.Sprintf("%d", v)
		} else {
			s += fmt.Sprintf("%d, ", v)
		}
	}
	return s + "]"
}

/*
func (c Coordinate) hash() uint32 {
	hasher := fnv.New32a()
	for _, num := range c {
		hasher.Write([]byte(strconv.Itoa(num)))
	}
	return hasher.Sum32()
}

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

// a cell has a state of type T
type Cell[T comparable] struct {
	State      T
	Coordinate Coordinate
}

func (c *Cell[T]) String() string {
	return c.Coordinate.String() + fmt.Sprintf(":%v", c.State)
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
func (bg *BaseGrid[T]) GetCellByHash(hash uint64) *Cell[T] {
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

func (bg *BaseGrid[T]) CheckIntegrity() {
	for _, shard := range bg.Cells.shards {
		for key, cell := range shard {
			if cell.Coordinate.hash() != key {
				panic(fmt.Sprintf("Integrity error: cell at key %v has coordinate %v\n", key, cell.Coordinate))
			}
		}
	}
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
//	fmt.Printf("coord: %v pre: %v\n", coord, neighborCoords)
	for _, co := range neighborCoords {
		cell := g.GetCell(co)
		neighbors = append(neighbors, cell)
	}
//	fmt.Printf("coord: %v post: %v\n", coord, neighbors)
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
