package base

type Grid struct {
	Width  int
	Height int
}

type Cells[T any] interface {
	// get the next layer from the passed layer index
	Next(layer int) []T
	// initialize the cells based on grid constraints
	Init()
	// Run the system from the initial conditions
	Run(initLayer []T)
}

// discrete cells can either be 1 or 0
type DiscreteCells struct {
	Grid  Grid
	Cells [][]bool
	Rules RuleSet[bool]
	NFunc Neighborhoods[bool]
}

func (dc *DiscreteCells) Next(layer int) []bool {
	neighborhoods := dc.NFunc(layer)

	next := make([]bool, dc.Grid.Width + 1)

	for i, nb := range neighborhoods {
		rule := dc.Rules.GetRule(nb)
		next[i] = rule(nb) 
	}
	return next
}

func (dc *DiscreteCells) Init() {
	dc.Cells = make([][]bool, dc.Grid.Height)
	for i := range dc.Cells {
		dc.Cells[i] = make([]bool, dc.Grid.Width)
	}
}

func (dc *DiscreteCells) Run(initialLayer []bool) {
	dc.Cells[0] = initialLayer
	// go up until the second to last row in grid
	for i := 0; i < dc.Grid.Height-1; i++ {
		n := dc.Next(i)
		dc.Cells[i+1] = n
	}
}
