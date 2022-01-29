package graphics

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
)

func Rectangle(x, y, width, height float32, texX, texY, texWidth, texHeight int, texture *gl.Texture) ([]mgl32.Vec2, []mgl32.Vec2, error) {
	if texX < 0 || texY < 0 || texX > texWidth || texY > texHeight {
		return nil, nil, fmt.Errorf("cannot create recangle, invalid texture coordinates")
	}

	vertices := make([]mgl32.Vec2, 6)
	texs := make([]mgl32.Vec2, 6)

	tX, tY := texture.PixToTex(texX, texY)
	tWidth := float32(texWidth) / float32(texture.Width)
	tHeight := float32(texHeight) / float32(texture.Height)

	/*
		We add height to the vertices in the complement of when we add height to the texs
		as the textures are referenced from the top left while the screen space is referenced from the bottom left
	*/

	// Upper left triangle

	// Top left
	vertices[0] = mgl32.Vec2{x, y + height}
	texs[0] = mgl32.Vec2{tX, tY}

	// Top right
	vertices[1] = mgl32.Vec2{x + width, y + height}
	texs[1] = mgl32.Vec2{tX + tWidth, tY}

	// Bottom Left
	vertices[2] = mgl32.Vec2{x, y}
	texs[2] = mgl32.Vec2{tX, tY + tHeight}

	// Lower right triangle

	// Bottom Left
	vertices[3] = mgl32.Vec2{x, y}
	texs[3] = mgl32.Vec2{tX, tY + tHeight}

	// Top right
	vertices[4] = mgl32.Vec2{x + width, y + height}
	texs[4] = mgl32.Vec2{tX + tWidth, tY}

	// Bottom right
	vertices[5] = mgl32.Vec2{x + width, y}
	texs[5] = mgl32.Vec2{tX + tWidth, tY + tHeight}

	return vertices, texs, nil
}

func Square(x, y, width float32, tX, tY, tWidth int, texture *gl.Texture) ([]mgl32.Vec2, []mgl32.Vec2, error) {
	return Rectangle(x, y, width, width, tX, tY, tWidth, tWidth, texture)
}
