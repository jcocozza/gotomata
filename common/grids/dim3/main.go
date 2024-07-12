package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWidth, screenHeight, "raylib [models] example - waving cubes")

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(30.0, 20.0, 30.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 70.0
	camera.Projection = rl.CameraPerspective

	numBloks := 15

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 5)
		for x := 0; x < numBloks; x++ {
			for y := 0; y < numBloks; y++ {
				for z := 0; z < numBloks; z++ {
					DrawCube(x, y, z, numBloks, 1)
				}
			}
		}
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func rotator(camera rl.Camera) {
	time := rl.GetTime()
	camTime := time * 0.3
	camera.Position.X = float32(math.Cos(camTime)) * 40
	camera.Position.Z = float32(math.Sin(camTime)) * 40
}

func DrawCube(x, y, z, numBloks, state int) {
	blockScale := float32((x + y + z)) / 30
	cubePos := rl.NewVector3(float32((x - numBloks/2)), float32((y - numBloks/2)), float32((z - numBloks/2)))
	cubeSize := (2.4) * blockScale

	var cubeColor rl.Color
	switch state {
	case 0:
		cubeColor = rl.White
	case 1:
		cubeColor = rl.Green
	case 2:
		cubeColor = rl.Yellow
	case 3:
		cubeColor = rl.Orange
	case 4:
		cubeColor = rl.Red
	default:
		cubeColor = rl.Gray // Default color for undefined states
	}

	rl.DrawCube(cubePos, cubeSize, cubeSize, cubeSize, cubeColor)
}
