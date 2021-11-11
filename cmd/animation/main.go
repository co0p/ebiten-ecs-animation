package main

import (
	"bytes"
	ecsexample "ebiten-ecs-animation"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:generate file2byteslice -package=main -input=spritesheet.png -output=spritesheet.go -var=Spritesheet

const screenWidth = 500
const screenHeight = 500

func main() {

	// loading assets
	frames := LoadSpritesheet(Spritesheet, 4, 250, 260)

	// entities
	registry := ecsexample.Registry{}

	e := registry.NewEntity()
	e.AddComponent(&ecsexample.AnimationComponent{
		Frames:            frames,
		CurrentFrameIndex: 0,
		Count:             0,
		AnimationSpeed:    0.125,
	})
	e.AddComponent(&ecsexample.SpriteComponent{Image: frames[0]})
	e.AddComponent(&ecsexample.TransformComponent{PosX: 100, PosY: 100})

	// systems
	animationSystem := ecsexample.AnimationSystem{Registry: &registry}
	spriteRenderSystem := ecsexample.SpriteRenderSystem{Registry: &registry}

	// the game
	example := AnimationExample{
		spriteRenderSystem: &spriteRenderSystem,
		animationSystem:    &animationSystem,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("animation example")
	ebiten.RunGame(&example) // omitted error handling
}

// LoadSpritesheet returns n sub images from the given input image
func LoadSpritesheet(input []byte, n int, width int, height int) []*ebiten.Image {
	var sprites []*ebiten.Image

	spritesheet, _, _ := image.Decode(bytes.NewReader(input))
	ebitenImage := ebiten.NewImageFromImage(spritesheet)

	for i := 0; i < n; i++ {
		dimensions := image.Rect(i*width, 0, (i+1)*width, height)
		sprite := ebitenImage.SubImage(dimensions).(*ebiten.Image)
		sprites = append(sprites, sprite)
	}

	return sprites
}
