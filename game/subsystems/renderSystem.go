package subsystems

import (
	"../components"
	"../ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderSystem struct {
	Em *ecs.EntityManager
}

func (system RenderSystem) Update() {
	for _, val := range system.Em.GetAllEntities() {
		positionComponent, isPosPresent := system.Em.GetComponent(val, ecs.COMPONENT_POSITION)

		if !isPosPresent {
			continue
		}

		pos := positionComponent.Data.(*components.Position)

		spriteComponent, isSpritePresent := system.Em.GetComponent(val, ecs.COMPONENT_SPRITE)
		if isSpritePresent {
			sprite := spriteComponent.Data.(*components.Sprite)
			rl.DrawTextureRec(*sprite.Texture, sprite.Rect, pos.Vector2, rl.White)
		}

		boxComponent, isBoxPresent := system.Em.GetComponent(val, ecs.COMPONENT_BOX)
		if isBoxPresent {
			box := boxComponent.Data.(*components.Box)
			rl.DrawRectangle(int32(pos.X+0.5), int32(pos.Y+0.5), int32(box.Width+0.5), int32(box.Height+0.5), box.Color)
		}

		circleComponent, isCirclePresent := system.Em.GetComponent(val, ecs.COMPONENT_CIRCLE)
		if isCirclePresent {
			circle := circleComponent.Data.(*components.Circle)
			rl.DrawCircle(int32(pos.X+0.5), int32(pos.Y+0.5), circle.Radius, circle.Color)
		}
	}
}
