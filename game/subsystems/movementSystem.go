package subsystems

import (
	"../components"
	"../ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MotionSystem struct {
	Em *ecs.EntityManager
}

func (system MotionSystem) Update() {
	dt := rl.GetFrameTime()

	for _, query := range system.Em.GetFilteredEntities([]ecs.ComponentType{ecs.COMPONENT_POSITION, ecs.COMPONENT_MOTION}) {
		position := query.Components[0].(*components.Position)
		motion := query.Components[1].(*components.Motion)

		position.X += motion.X * dt
		position.Y += motion.Y * dt
	}
}
