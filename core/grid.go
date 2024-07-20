package core

import "fmt"

// an n-dimensional coordinate on the grid
type Coordinate []int

func (c Coordinate) hash() uint64 {
	var hash uint64 = 14695981039346656037 // FNV offset basis
	for i, num := range c {
		// Incorporate the index to differentiate [1,2,3] from [3,2,1]
		hash ^= uint64(i)
		hash *= 1099511628211
		hash ^= uint64(num)
		hash *= 1099511628211
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

// a cell has a state of type T
type Cell[T comparable] struct {
	State      T
	Coordinate Coordinate
}

// useful for debugging
func (c *Cell[T]) String() string {
	return c.Coordinate.String() + fmt.Sprintf(":%v", c.State)
}

type BaseGrid[T comparable] struct {
	// a list of dimension sizes
	// i.e for a 4x4 grid, dimensions = []int{4,4}
	//
	// note that you do not necessarily have to specify these
	// instead, max dims can be calculated via ComputeMaxDims()
	// these can be used to determine grid sizes
	//
	// however, if you want to use a finite grid size and handle boundary conditions
	// specifing the max dimensions will likely be better
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

// set a list of coordinates to the passed state
func (bg *BaseGrid[T]) SetConfig(cfg []Coordinate, state T) {
	for _, coord := range cfg {
		bg.SetCell(state, coord)
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

// pass in a list of dimensions
//
// if the list is nil, use the dimensions specified in BaseGrid
//
// this will probably not work with hexagonal grid systems because of negative coordinates
func (bg *BaseGrid[T]) AllCoordinates(dims []int) []Coordinate {
	if dims == nil {
		dims = bg.Dimensions
	}
	totalCoords := 1
	for _, dim := range dims {
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

// ensure that each cell coordinate's hash is assigned properly to the key
func (bg *BaseGrid[T]) CheckIntegrity() {
	for _, shard := range bg.Cells.shards {
		for key, cell := range shard {
			if cell.Coordinate.hash() != key {
				panic(fmt.Sprintf("Integrity error: cell at key %v has coordinate %v\n", key, cell.Coordinate))
			}
		}
	}
}

// get the maximum value along each axis in the grid (in parallel)
//
// can be used to dynamically size output drawings
func (bg *BaseGrid[T]) ComputeMaxDims() []int {
	maxForEachShard := make([][]int, len(bg.Cells.shards))
	var processer = func(shard int, cells map[uint64]*Cell[T]) {
		var initCell *Cell[T]
		for key := range cells {
			initCell = bg.GetCellByHash(key)
			if initCell != nil {
				break
			}
		}

		// handle case when shard is empty
		if initCell == nil {
			maxForEachShard[shard] = nil
			return
		}

		max := make([]int, len(initCell.Coordinate))
		copy(max, initCell.Coordinate)

		for _, cell := range cells {
			for i, loc := range cell.Coordinate {
				if max[i] < loc {
					max[i] = loc
				}
			}
		}

		maxForEachShard[shard] = max
	}
	bg.Cells.ProcessShard(processer)

	var s int
	for _, m := range maxForEachShard {
		if m != nil {
			s = len(m)
		}
	}
	// collect all the maxes in each shard to a total max
	overallMax := make([]int, s)
	for _, maximum := range maxForEachShard {
		if maximum == nil {
			continue
		}
		for i, m := range maximum {
			if overallMax[i] < m {
				overallMax[i] = m
			}
		}
	}
	return overallMax
}

// return all the neighbors of a coordinate
type GetNeighborsFunc func(coord Coordinate) []Coordinate

// a grid is a set of organized cells with a "geometry" (i.e. a way to get neighbors)
type Grid[T comparable] struct {
	*BaseGrid[T]
	GetNeighborCoordinates GetNeighborsFunc
}

// Use the neighbor geometry function to get the cell neighbors
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
