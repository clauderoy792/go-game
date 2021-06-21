package subsystems

import (
	"../components"
	"../ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicSystem struct {
	Em *ecs.EntityManager
}

type collisionPair struct {
	entityA    ecs.EntityID
	entityB    ecs.EntityID
	overlapDst float32
}

func (system PhysicSystem) Update() {
	dt := rl.GetFrameTime()

	/*
		entities := system.Em.GetAllEntitiesOfComponentType(ecs.COMPONENT_CIRCLE_COLLIDER)
		entitiesCount := len(entities)

		collisions := make([]collisionPair, 0)

		// 1. Calculate Collisions Pairs
		for i := 0; i < entitiesCount; i++ {
			posCompA, present := system.Em.GetComponent(entities[i], ecs.COMPONENT_POSITION)

			if !present {
				continue
			}

			colliderCompA, _ := system.Em.GetComponent(entities[i], ecs.COMPONENT_CIRCLE_COLLIDER)
			colliderA := colliderCompA.Data.(*components.CircleCollider)
			posA := posCompA.Data.(*components.Position)

			for j := i + 1; j < entitiesCount; j++ {
				posCompB, present := system.Em.GetComponent(entities[j], ecs.COMPONENT_POSITION)

				if !present {
					continue
				}

				colliderCompB, _ := system.Em.GetComponent(entities[j], ecs.COMPONENT_CIRCLE_COLLIDER)
				colliderB := colliderCompB.Data.(*components.CircleCollider)
				posB := posCompB.Data.(*components.Position)

				dst := raymath.Vector2Distance(posA.Vector2, posB.Vector2)
				if dst < colliderA.Radius+colliderB.Radius {
					collisions = append(collisions, collisionPair{entityA: entities[i], entityB: entities[j], overlapDst: (colliderA.Radius + colliderB.Radius) - dst})
				}
			}
		}
	*/

	// 2. Resolve Collisions
	/*for _, collision := range collisions {
		motionA, aHasMotion := system.Em.GetComponent(collision.entityA, ecs.COMPONENT_MOTION)
		motionB, bHasMotion := system.Em.GetComponent(collision.entityB, ecs.COMPONENT_MOTION)

		if aHasMotion && bHasMotion {

		} else if aHasMotion {

		} else if bHasMotion {

		}
	}*/

	for _, query := range system.Em.GetFilteredEntities([]ecs.ComponentType{ecs.COMPONENT_PHYSIC_BODY, ecs.COMPONENT_MOTION}) {
		physicBody := query.Components[0].(*components.PhysicBody)
		motion := query.Components[1].(*components.Motion)

		// Apply Gravity
		motion.Y += 9.8 * physicBody.Mass * physicBody.GravityMultiplier * dt // TODO: * PixelPerUnits
	}
}
