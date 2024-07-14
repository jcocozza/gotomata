package grids

import (
	"math/rand"

	"github.com/jcocozza/gotomata/core"
)

func Dim3Grid[T comparable](x, y, z int, defaultState T, neighborsFunc core.GetNeighborsFunc) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions: []int{x, y, z},
			Cells: core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: neighborsFunc,
	}
}

func InitSphere(width, height, depth, aliveState int, c *core.CellularAutomata[int]) {
	radius := 5// Define the radius of the sphere
	center := core.Coordinate{width / 2, height / 2, depth / 2}

	for x := center[0] - radius; x <= center[0] + radius; x++ {
		for y := center[1] - radius; y <= center[1] + radius; y++ {
			for z := center[2] - radius; z <= center[2] + radius; z++ {
				if (x-center[0])*(x-center[0])+(y-center[1])*(y-center[1])+(z-center[2])*(z-center[2]) <= radius*radius {
					p := rand.Float64()
					if p < .3 {
						c.Grid.SetCell(aliveState, core.Coordinate{x, y, z})
					}
				}
			}
		}
	}
}

func InitCube(width, height, depth, aliveState int, c *core.CellularAutomata[int]) {
	center := core.Coordinate{width / 2, height / 2, depth / 2}

	sideLength := 5// Define the side length of the cube

	halfSide := sideLength / 2
	for x := center[0] - halfSide; x <= center[0]+halfSide; x++ {
		for y := center[1] - halfSide; y <= center[1]+halfSide; y++ {
			for z := center[2] - halfSide; z <= center[2]+halfSide; z++ {
				p := rand.Float64()
				if p < .55 {
					c.Grid.SetCell(aliveState, core.Coordinate{x, y, z})
				}
			}
		}
	}
}
