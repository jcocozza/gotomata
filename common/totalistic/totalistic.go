package totalistic

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
	"github.com/jcocozza/gotomata/utils"
)

var Neighbors = func(coord core.Coordinate) []core.Coordinate {
	left := core.Coordinate{(coord[0] - 1)}
	right := core.Coordinate{(coord[0] + 1)}
	return []core.Coordinate{left, right}
}

func grid(length int) *core.Grid[int] {
	return grids.Dim1Grid[int](length, 0, Neighbors)
}

func base10ToBase3(num int) int {
	if num == 0 {
		return 0
	}

	var base3 int
	multiplier := 1

	for num > 0 {
		remainder := num % 3
		base3 += remainder * multiplier
		num /= 3
		multiplier *= 10
	}

	return base3
}

// GetNthDigit returns the nth digit (0-indexed from the right) of an integer.
func getNthDigit(num, n int) int {
	if num < 0 {
		num = -num // Handle negative numbers
	}

	// Calculate the divisor to extract the nth digit.
	divisor := int(math.Pow(10, float64(n)))

	// Check if the divisor is larger than the number.
	// if it is, then we just need to 0 pad
	if divisor > num {
		return 0
	}

	// Get the nth digit.
	digit := (num / divisor) % 10

	return digit
}

// totalistic rule numbers correspond to a sequence of base 3 digits
//
// we consider the average of the 3 states in the neighborhood.
// there are 7 possible averages: 2, 5/3, 4/3, 1, 2/3, 1/3, 0
//
// thus we have 3^7 = 2187 total possible rules
//
// now we can encode a rule as a base 3 number where each average is mapped to a state
// for example:
// 2   -> 1
// 5/3 -> 2
// 4/3 -> 0
// 1   -> 2
// 2/3 -> 1
// 1/3 -> 0
// 0   -> 2
//
// this rule, in base 3 is: 1202102, so the base 10 rule number is: 1280
func TotalisticRuleSet(ruleNumber int) core.RuleSet[int] {
	var enumMap = map[float64]int{
		2.0:       6,
		5.0 / 3.0: 5,
		4.0 / 3.0: 4,
		1.0:       3,
		2.0 / 3.0: 2,
		1.0 / 3.0: 1,
		0.0:       0,
	}
	b3 := base10ToBase3(ruleNumber)

	ruleset := func(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {
		left := neighbors[0].State
		center := cell.State
		right := neighbors[1].State
		avg := float64(left+center+right) / 3

		pos := enumMap[avg]
		if pp, exists := enumMap[avg]; !exists {
			fmt.Println(pp)
			panic("not found")
		}
		state := getNthDigit(b3, pos)

		return &core.Cell[int]{
			State:      state,
			Coordinate: cell.Coordinate,
		}
	}
	return ruleset
}

func TotalisticCellularAutomata(ruleNum, length, steps int) *core.CellularAutomata[int] {
	g := grid(length)
	return &core.CellularAutomata[int] {
		Grid: g,
		RuleSet: TotalisticRuleSet(ruleNum),
		Steps: steps,
	}
}

func SetCenterConfig(length int) []core.Coordinate{
	return []core.Coordinate{{length / 2}}
}

func MainTotalistic(ruleNum, length, steps, scale int, initConfig []core.Coordinate) {
	ca := TotalisticCellularAutomata(ruleNum, length, steps)
	for _, coord := range initConfig {
		ca.Grid.SetCell(1, coord)
	}
	img := image.NewGray(image.Rect(0,0, length * scale, steps * scale))
	AddTotalisticToImage(ca, img, 0, scale, length)
	for i := 0; i < steps; i++ {
		ca.Stepp()
		if i < steps {
			AddTotalisticToImage(ca, img, i+1, scale, length)
		}
	}
	utils.WritePNG(img, "_images/totalistic.png")
}

func stateToColor(state int) color.Color {
	switch state {
	case 0:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case 1:
		return color.RGBA{R: 128, G: 128, B: 128, A: 255}
	case 2:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	default:
		panic("no proper state")
	}
}

func AddTotalisticToImage(ca *core.CellularAutomata[int], img *image.Gray, stepNum, scale, length int) {
	coords := ca.Grid.AllCoordinates(nil)
	for i, coord := range coords {
		cell := ca.Grid.GetCell(coord)
		color := stateToColor(cell.State)
		for dx := 0; dx < scale; dx++ {
			for dy := 0; dy < scale; dy++ {
				img.Set((i*scale)+dx, (stepNum*scale) + dy, color)
			}
		}
	}
}


