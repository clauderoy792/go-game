package game

import (
	"./components"
	"./ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewMotionComponent(x, y float32) *ecs.Component {
	return ecs.NewComponent(&components.Motion{Vector2: rl.Vector2{X: x, Y: y}}, ecs.COMPONENT_MOTION)
}

func NewPositionComponent(x, y float32) *ecs.Component {
	return ecs.NewComponent(&components.Position{Vector2: rl.Vector2{X: x, Y: y}}, ecs.COMPONENT_POSITION)
}

func NewSpriteComponent(texture *rl.Texture2D, rect rl.Rectangle) *ecs.Component {
	return ecs.NewComponent(&components.Sprite{Texture: texture, Rect: rect}, ecs.COMPONENT_SPRITE)
}

func NewUserInputAIComponent(speed float32) *ecs.Component {
	return ecs.NewComponent(&components.UserInputAI{Speed: speed}, ecs.COMPONENT_USER_INPUT_AI)
}

func NewCircleComponent(radius float32, color rl.Color) *ecs.Component {
	return ecs.NewComponent(&components.Circle{Radius: radius, Color: color}, ecs.COMPONENT_CIRCLE)
}

func NewBoxComponent(width, height float32, color rl.Color) *ecs.Component {
	return ecs.NewComponent(&components.Box{Width: width, Height: height, Color: color}, ecs.COMPONENT_BOX)
}

func NewPhysicBodyComponent(mass, gravityMultiplier float32) *ecs.Component {
	return ecs.NewComponent(&components.PhysicBody{Mass: mass, GravityMultiplier: gravityMultiplier}, ecs.COMPONENT_PHYSIC_BODY)
}

func NewCircleColliderComponent(radius float32) *ecs.Component {
	return ecs.NewComponent(&components.CircleCollider{Radius: radius}, ecs.COMPONENT_CIRCLE_COLLIDER)
}
