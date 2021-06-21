package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Texture *rl.Texture2D
	Rect    rl.Rectangle
}
