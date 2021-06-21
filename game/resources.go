package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var textures = make(map[string]rl.Texture2D)

func loadTexture(fileName string) *rl.Texture2D {
	tex, exists := textures[fileName]
	if !exists {
		tex = rl.LoadTexture(fileName)
		textures[fileName] = tex
	}

	return &(tex)
}

func unloadTexture(fileName string) {
	_, exists := textures[fileName]
	if exists {
		delete(textures, fileName)
	}
}
