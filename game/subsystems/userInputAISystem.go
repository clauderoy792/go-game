package subsystems

import (
	"../components"
	"../ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
)

type UserInputAISystem struct {
	Em *ecs.EntityManager
}

func (system UserInputAISystem) Update() {
	direction := rl.Vector2{}

	if rl.IsKeyDown(rl.KeyW) {
		direction.Y++
	}
	if rl.IsKeyDown(rl.KeyS) {
		direction.Y--
	}
	if rl.IsKeyDown(rl.KeyD) {
		direction.X++
	}
	if rl.IsKeyDown(rl.KeyA) {
		direction.X--
	}

	if direction.X != 0 || direction.Y != 0 {
		raymath.Vector2Normalize(&direction)
	}

	for _, query := range system.Em.GetFilteredEntities([]ecs.ComponentType{ecs.COMPONENT_USER_INPUT_AI, ecs.COMPONENT_MOTION}) {
		userController := query.Components[0].(*components.UserInputAI)
		motion := query.Components[1].(*components.Motion)

		motion.X = direction.X * userController.Speed
		motion.Y = -direction.Y * userController.Speed
	}
}
