package game

import (
	"math/rand"

	"./ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func newBasicEntity(em *ecs.EntityManager, x, y float32) ecs.EntityID {
	entity := em.CreateEntity()
	em.AddComponent(entity, NewPositionComponent(x, y))
	return entity
}

func newPlayer(em *ecs.EntityManager, x, y float32) ecs.EntityID {
	entity := newBasicEntity(em, x, y)
	em.AddComponent(entity, NewSpriteComponent(loadTexture("./assets/tilemaps/dungeon_tileset.png"), rl.Rectangle{X: 128, Y: 64, Width: 16, Height: 32}))
	em.AddComponent(entity, NewMotionComponent(0, 0))
	em.AddComponent(entity, NewUserInputAIComponent(60))

	return entity
}

func newBox(em *ecs.EntityManager, x, y, width, height float32, color rl.Color) ecs.EntityID {
	entity := newBasicEntity(em, x, y)
	em.AddComponent(entity, NewBoxComponent(width, height, color))
	return entity
}

func newCircle(em *ecs.EntityManager, x, y, radius float32, color rl.Color) ecs.EntityID {
	entity := newBasicEntity(em, x, y)
	em.AddComponent(entity, NewCircleComponent(radius, color))
	return entity
}

func newRandomBall(em *ecs.EntityManager, x, y, radius float32) ecs.EntityID {
	tint := rl.NewColor(uint8(rand.Float32()*255), uint8(rand.Float32()*255), uint8(rand.Float32()*255), 255)
	entity := newCircle(em, x, y, rand.Float32()*30, tint)
	em.AddComponent(entity, NewMotionComponent(rand.Float32()*100-50, rand.Float32()*100-50))
	em.AddComponent(entity, NewPhysicBodyComponent(radius*2, 1))
	return entity
}
