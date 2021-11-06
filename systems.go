package ecsexample

import "github.com/hajimehoshi/ebiten/v2"

type SpriteRenderSystem struct {
	Registry *Registry
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {

	for _, e := range s.Registry.Query(TransformType, SpriteType) {

		position := e.GetComponent(TransformType).(*TransformComponent)
		sprite := e.GetComponent(SpriteType).(*SpriteComponent)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.PosX, position.PosY)
		screen.DrawImage(sprite.Image, op)
	}
}
