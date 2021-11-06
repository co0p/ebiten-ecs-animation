package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// ComponentType defines the supported component types in a user readable format
type ComponentType string

// ComponentTyper returns the type of a component
type ComponentTyper interface{ Type() ComponentType }

type Entity struct {
	components map[ComponentType]ComponentTyper
}

// NewEntity returns a new initialized entity
func NewEntity() Entity { return Entity{components: make(map[ComponentType]ComponentTyper)} }

// HasComponent returns true if the entity has the component of type cType associated
func (e *Entity) HasComponent(cType ComponentType) bool {
	_, ok := e.components[cType]
	return ok
}

// GetComponent returns the component, panics if not found
func (e *Entity) GetComponent(cType ComponentType) ComponentTyper {
	if c, ok := e.components[cType]; !ok {
		panic(fmt.Sprintf("expected entity to have component of type %s attached", cType))
	} else {
		return c
	}
}

// AddComponent adds a component to the entity, panics if a component of the same type already exists
func (e *Entity) AddComponent(c ComponentTyper) {
	if e.HasComponent(c.Type()) {
		panic(fmt.Sprintf("entity already has component of type %s attached", c))
	}
	e.components[c.Type()] = c
}

const TransformType ComponentType = "TRANSFORM"

type TransformComponent struct {
	posX, posY float64
}

func (t *TransformComponent) Type() ComponentType { return TransformType }

const SpriteType ComponentType = "SPRITE"

type SpriteComponent struct {
	image *ebiten.Image
}

func (t *SpriteComponent) Type() ComponentType { return SpriteType }
