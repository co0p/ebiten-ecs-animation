package main

import "github.com/hajimehoshi/ebiten/v2"

type SpriteRenderSystem struct {
	registry *Registry
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {

	for _, e := range s.registry.Query(TransformType, SpriteType) {

		position := e.GetComponent(TransformType).(*TransformComponent)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.posX, position.posY)

		sprite := e.GetComponent(SpriteType).(*SpriteComponent)
		screen.DrawImage(sprite.image, op)
	}
}
