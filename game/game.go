package game

import (
	"fmt"

	"./ecs"
	"./subsystems"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type Game struct {
	RenderWidth  int32
	RenderHeight int32
	RenderScale  float32

	entityManager *ecs.EntityManager
	playerEntity  ecs.EntityID
	debugTexture  rl.Texture2D
	systems       []ecs.System

	tilemap *tiled.Map
}

func NewGame() Game {
	g := Game{}
	return g
}

func (g *Game) Initialize() {
	g.entityManager = ecs.NewEntityManager()

	g.playerEntity = newPlayer(g.entityManager, float32(g.RenderWidth)/2, float32(g.RenderHeight)/2)

	/*
		for i := 0; i < 10000; i++ {
			newRandomBall(g.entityManager, rand.Float32()*float32(g.RenderWidth), rand.Float32()*float32(g.RenderHeight), 10)
		}
	*/

	g.systems = []ecs.System{
		subsystems.UserInputAISystem{Em: g.entityManager},
		subsystems.MotionSystem{Em: g.entityManager},
		// subsystems.PhysicSystem{Em: g.entityManager},
		subsystems.RenderSystem{Em: g.entityManager},
	}

	gameMap, _ := tiled.LoadFromFile("./assets/tilemaps/test_map.tmx")
	g.tilemap = gameMap
}

// Tick is called by the main program when we process a new frame
func (g *Game) Tick() {
	DrawTilemap(g.tilemap)

	for _, system := range g.systems {
		system.Update()
	}

	// FPS Counter
	rl.DrawText(fmt.Sprintf("%v", 1/rl.GetFrameTime()), 5, 5, 10, rl.White)
}

func DrawTilemap(tilemap *tiled.Map) {
	tileset := loadTexture("./assets/tilemaps/main_tilemap.png")
	tileW := float32(tilemap.TileWidth)
	tileH := float32(tilemap.TileWidth)

	for _, layer := range tilemap.Layers {
		for tileIndex, tile := range layer.Tiles {
			id := int32(tile.ID)

			if id == 0 {
				continue
			}

			tileSrcX := float32(id%(tileset.Width/int32(tileW))) * tileW
			tileSrcY := float32(id/(tileset.Width/int32(tileW))) * tileH
			tileDstX := float32(tileIndex%tilemap.Width) * tileW
			tileDstY := float32(tileIndex/tilemap.Width) * tileH

			source := rl.NewRectangle(tileSrcX, tileSrcY, tileW, tileH)
			destination := rl.NewRectangle(tileDstX, tileDstY, tileW, tileH)
			rotation := float32(0)

			if tile.DiagonalFlip {
				rotation = 270
				source.Height *= -1
				destination.Y += destination.Height
			}

			if tile.HorizontalFlip {
				source.Width *= -1
			}

			if tile.VerticalFlip {
				source.Height *= -1
			}

			rl.DrawTexturePro(*tileset, source, destination, rl.Vector2{}, rotation, rl.White)
		}
	}
}
