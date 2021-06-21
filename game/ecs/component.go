package ecs

type ComponentType uint8

type Component struct {
	ComponentType ComponentType
	Data          interface{}
}

const (
	COMPONENT_POSITION ComponentType = iota
	COMPONENT_MOTION
	COMPONENT_SPRITE
	COMPONENT_USER_INPUT_AI
	COMPONENT_CIRCLE
	COMPONENT_BOX
	COMPONENT_PHYSIC_BODY
	COMPONENT_CIRCLE_COLLIDER
)

func NewComponent(data interface{}, componentType ComponentType) *Component {
	return &Component{componentType, data}
}
