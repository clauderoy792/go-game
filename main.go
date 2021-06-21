package main

import (
	"./game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := game.Game{RenderWidth: 640, RenderHeight: 360, RenderScale: 2}

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(float32(game.RenderWidth)*game.RenderScale), int32(float32(game.RenderHeight)*game.RenderScale), "Test")
	rl.SetTargetFPS(60)

	defer rl.CloseWindow()

	game.Initialize()

	renderTexture := rl.LoadRenderTexture(game.RenderWidth, game.RenderHeight)
	rl.SetTextureFilter(renderTexture.Texture, rl.FilterPoint)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginTextureMode(renderTexture)

		game.Tick()

		rl.EndTextureMode()

		// Calculate Letterboxing if the aspect ratio is not the same as the native render size
		windowWidth := float32(rl.GetScreenWidth())
		windowHeight := float32(rl.GetScreenHeight())
		targetWidth := windowWidth
		targetHeight := windowHeight

		if float32(game.RenderWidth)/float32(game.RenderHeight) > windowWidth/windowHeight {
			targetHeight = windowWidth * float32(game.RenderHeight) / float32(game.RenderWidth)
		} else {
			targetWidth = windowHeight * float32(game.RenderWidth) / float32(game.RenderHeight)
		}

		rl.DrawTexturePro(
			renderTexture.Texture,
			rl.NewRectangle(0, 0, float32(game.RenderWidth), float32(-game.RenderHeight)),
			rl.NewRectangle((windowWidth-targetWidth)/2, (windowHeight-targetHeight)/2, targetWidth, targetHeight),
			rl.NewVector2(0, 0),
			0,
			rl.White)

		rl.EndDrawing()
	}
}
