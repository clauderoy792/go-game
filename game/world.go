package game

import "./ecs"

type world struct {
}

type worldChunk struct {
	entities []ecs.EntityID
}
