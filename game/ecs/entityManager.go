package ecs

type EntityID uint64

type EntityManager struct {
	lowestUnassignedID EntityID
	entities           []EntityID
	componentStores    map[ComponentType]map[EntityID]Component
}

type QueryResult struct {
	Entity     EntityID
	Components []interface{}
}

func NewEntityManager() *EntityManager {
	return &EntityManager{1, make([]EntityID, 0), make(map[ComponentType]map[EntityID]Component)}
}

func (e *EntityManager) GetFilteredEntities(query []ComponentType) []QueryResult {
	results := make([]QueryResult, 0)

	for _, entity := range e.entities {
		result := QueryResult{Entity: entity, Components: make([]interface{}, 0)}

		for _, componentType := range query {
			if component, isPresent := e.componentStores[componentType][entity]; isPresent {
				result.Components = append(result.Components, component.Data)
			}
		}

		if len(result.Components) == len(query) {
			results = append(results, result)
		}
	}

	return results
}

func (e *EntityManager) GetComponent(entity EntityID, componentType ComponentType) (*Component, bool) {
	val, _ := e.componentStores[componentType]
	value, isPresent := val[entity]
	return &value, isPresent

}

func (e *EntityManager) GetAllComponentsOfType(componentType ComponentType) []*Component {
	val, _ := e.componentStores[componentType]
	allComponents := make([]*Component, 0)
	for _, value := range val {
		allComponents = append(allComponents, &value)

	}
	return allComponents

}

func (e *EntityManager) GetAllEntitiesOfComponentType(componentType ComponentType) []EntityID {
	val, _ := e.componentStores[componentType]
	allEntities := make([]EntityID, 0)
	for key := range val {
		allEntities = append(allEntities, key)
	}
	return allEntities

}

func (e *EntityManager) AddComponent(entity EntityID, component *Component) {
	_, isPresent := e.componentStores[component.ComponentType]
	if isPresent {
		e.componentStores[component.ComponentType][entity] = *component
	} else {
		e.componentStores[component.ComponentType] = map[EntityID]Component{entity: *component}
	}
}

func (e *EntityManager) CreateEntity() EntityID {
	newEntityID := e.generateEntityID()
	e.entities = append(e.entities, newEntityID)
	return newEntityID

}

func (e *EntityManager) DeleteEntity(entity EntityID) {
	e.entities = append(e.entities[:entity], e.entities[entity+1:]...)
}

func (e *EntityManager) GetAllEntities() []EntityID {
	return e.entities
}

func (e *EntityManager) generateEntityID() EntityID {
	e.lowestUnassignedID++
	return e.lowestUnassignedID
}
