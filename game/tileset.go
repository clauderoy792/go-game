package game

import (
	"./components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileSet struct {
	texture *rl.Texture2D
	sprites map[string]*components.Sprite
}

/*
func LoadTileSet(textureFileName, tilesFileName string) TileSet {

	tileSet := TileSet{}
}*/
