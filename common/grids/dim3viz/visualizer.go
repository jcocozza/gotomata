package dim3viz

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jcocozza/gotomata/core"
)

var (
	screenWidth         = int32(1280)   // Framebuffer width
	screenHeight        = int32(800)   // Framebuffer height
	fps                 = 60           // Frames per second

	speed               = 30           // Speed of jump animation
	groups              = 2            // Count of separate groups jumping around
	amp                 = float32(10)  // Maximum amplitude of jump
	variance            = float32(0.8) // Global variance in jump height
	loop                = float32(0.0) // Individual cube's computed loop timer
	textPositionY int32 = 300
	framesCounter       = 0
)

func computeTranslations[T comparable](ca *core.CellularAutomata[T]) []rl.Matrix {
	translations := make([]rl.Matrix, ca.Grid.Cells.Size())

	shiftX := float32(-ca.Grid.Dimensions[0])/2
	shiftY := float32(-ca.Grid.Dimensions[1])/2
	shiftZ := float32(-ca.Grid.Dimensions[2])/2

	var cnt int
	for _, key := range ca.Grid.Cells.GetAllKeys() {
		cell := ca.Grid.GetCellByHash(key)
		x := float32(cell.Coordinate[0]) + shiftX
		y := float32(cell.Coordinate[1]) + shiftY
		z := float32(cell.Coordinate[2]) + shiftZ
		translations[cnt] = rl.MatrixTranslate(x,y,z)
		cnt++
	}
	return translations
}

func Visualizer[T comparable](ca *core.CellularAutomata[T]) {
	rl.SetConfigFlags(rl.FlagMsaa4xHint) // Enable Multi Sampling Anti Aliasing 4x (if available)
	rl.InitWindow(screenWidth, screenHeight, "gotomata")


	// Define the camera to look into our 3d world
	camera := rl.Camera{
		Position:   rl.NewVector3(-125.0, 125.0, -125.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	cube := rl.GenMeshCube(1.0, 1.0, 1.0)

	shader := rl.LoadShader("/Users/josephcocozza/Repositories/gotomata/common/grids/dim3viz/glsl330/base_lighting_instanced.vs", "/Users/josephcocozza/Repositories/gotomata/common/grids/dim3viz/glsl330/lighting.fs")
	shader.UpdateLocation(rl.ShaderLocMatrixMvp, rl.GetShaderLocation(shader, "mvp"))
	shader.UpdateLocation(rl.ShaderLocVectorView, rl.GetShaderLocation(shader, "viewPos"))
	shader.UpdateLocation(rl.ShaderLocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))

	// ambient light level
	ambientLoc := rl.GetShaderLocation(shader, "ambient")
	rl.SetShaderValue(shader, ambientLoc, []float32{0.2, 0.2, 0.2, 1.0}, rl.ShaderUniformVec4)
	NewLight(LightTypeDirectional, rl.NewVector3(50.0, 50.0, 0.0), rl.Vector3Zero(), rl.White, shader)

	material := rl.LoadMaterialDefault()
	material.Shader = shader
	mmap := material.GetMap(rl.MapDiffuse)
	mmap.Color = rl.Red

	rl.SetTargetFPS(int32(fps))
	var step int
	for !rl.WindowShouldClose() {

		rl.SetShaderValue(shader, shader.GetLocation(rl.ShaderLocVectorView),
			[]float32{camera.Position.X, camera.Position.Y, camera.Position.Z}, rl.ShaderUniformVec3)

		rl.UpdateCamera(&camera, rl.CameraOrbital) // Update camera with orbital camera mode
		//rl.UpdateCamera(&camera, rl.CameraFree)


		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			rl.BeginMode3D(camera)

			size := ca.Grid.Cells.Size()
			translations := computeTranslations(ca)
			if len(translations) > 0 {
				rl.DrawMeshInstanced(cube, material, translations, size)
			}
			rl.EndMode3D()
			rl.DrawFPS(10,10)
			rl.DrawText(fmt.Sprintf("Step: %d", step), 490, 10, 20, rl.Maroon)
				if step == 0 {
					time.Sleep(2 * time.Second)
				}
			if step < ca.Steps {
				ca.Stepp()
				step++
			}
		}
		rl.EndDrawing()
		time.Sleep(100 * time.Millisecond)
	}
	rl.CloseWindow()
}
