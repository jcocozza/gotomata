package crystal3d

import (
	"math"
	//"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jcocozza/gotomata/core"
)

func Crystal(width, height, depth, steps int) *core.CellularAutomata[bool] {
	grid := CrystalGrid(width, height, depth)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: CrystalGrowth,
		Steps:   steps,
	}
}

func distanceToColor(distance float64) rl.Color {

	//r := uint8((math.Sin(distance*0.2+0.5*math.Pi) + 1) * 127.5)
	//g := uint8((math.Sin(distance*0.2+1.5*math.Pi) + 1) * 127.5)
	//b := uint8((math.Sin(distance*0.2+2.5*math.Pi) + 1) * 127.5)

	r := uint8((math.Sin(distance*0.1) + 1) * 127.5)
	g := uint8((math.Sin(distance*0.1+2*math.Pi/3) + 1) * 127.5)
	b := uint8((math.Sin(distance*0.1+4*math.Pi/3) + 1) * 127.5)
	return rl.NewColor(r, g, b, 255)
}

func DrawCubeCrystal(x, y, z int) {
	blockScale := float32(0.25)

	cubePos := rl.NewVector3(float32(x)*blockScale, float32(y)*blockScale, float32(z)*blockScale)

	distance := math.Sqrt(float64(x*x + y*y + z*z))
	color := distanceToColor(distance)

	rl.DrawCubeWires(cubePos, blockScale, blockScale, blockScale, rl.Black)
	rl.DrawCube(cubePos, blockScale, blockScale, blockScale, color)
}

func VisualizeCrystal(crystal *core.CellularAutomata[bool]) {
	screenWidth, screenHeight := int32(1280), int32(900)

	shiftX, shiftY, shiftZ := -crystal.Grid.Dimensions[0]/2, -crystal.Grid.Dimensions[1]/2, -crystal.Grid.Dimensions[2]/2

	rl.InitWindow(screenWidth, screenHeight, "3D Crystal")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(10.0, 05.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 100 //70.0
	camera.Projection = rl.CameraPerspective

	rl.SetTargetFPS(60)

	var count int
	for !rl.WindowShouldClose() && count < crystal.Steps {

		// rotation
		//t := rl.GetTime()
		//camTime := t * 0.3
		////camera.Position.X = float32(math.Cos(camTime)) * 40
		////camera.Position.Z = float32(math.Sin(camTime)) * 40
		//targetX := float32(math.Cos(camTime)) * 40
		//targetZ := float32(math.Sin(camTime)) * 40
		//camera.Position.X += (targetX - camera.Position.X) * 0.1
		//camera.Position.Z += (targetZ - camera.Position.Z) * 0.1

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 5)

		crystal.Stepp()
		for _, key := range crystal.Grid.Cells.GetAllKeys() {
			cell := crystal.Grid.GetCellByHash(key)
			x, y, z := cell.Coordinate[0], cell.Coordinate[1], cell.Coordinate[2]
			DrawCubeCrystal(x+shiftX, y+shiftY, z+shiftZ)
		}

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
		count++
		//time.Sleep(500 * time.Millisecond) // * time.Second)
	}
	rl.CloseWindow()
}
