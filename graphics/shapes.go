package graphics

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
)

func Rectangle(x, y, width, height, texX, texY, texWidth, texHeight int, texture *gl.Texture, window *gl.Window) ([]mgl32.Vec2, []mgl32.Vec2, error) {
	if texX < 0 || texY < 0 || texX > texWidth || texY > texHeight {
		return nil, nil, fmt.Errorf("cannot create recangle, invalid texture coordinates")
	}

	vertices := make([]mgl32.Vec2, 6)
	texs := make([]mgl32.Vec2, 6)

	sX, sY := window.PixToScreen(x, y)
	sWidth := 2 * float32(width) / window.Width
	sHeight := 2 * float32(height) / window.Height
	tX, tY := texture.PixToTex(texX, texY)
	tWidth := float32(texWidth) / float32(texture.Width)
	tHeight := float32(texHeight) / float32(texture.Height)

	/*
		We add height to the vertices in the complement of when we add height to the texs
		as the textures are referenced from the top left while the screen space is referenced from the bottom left
	*/

	// Upper left triangle

	// Top left
	vertices[0] = mgl32.Vec2{sX, sY + sHeight}
	texs[0] = mgl32.Vec2{tX, tY}

	// Top right
	vertices[1] = mgl32.Vec2{sX + sWidth, sY + sHeight}
	texs[1] = mgl32.Vec2{tX + tWidth, tY}

	// Bottom Left
	vertices[2] = mgl32.Vec2{sX, sY}
	texs[2] = mgl32.Vec2{tX, tY + tHeight}

	// Lower right triangle

	// Bottom Left
	vertices[3] = mgl32.Vec2{sX, sY}
	texs[3] = mgl32.Vec2{tX, tY + tHeight}

	// Top right
	vertices[4] = mgl32.Vec2{sX + sWidth, sY + sHeight}
	texs[4] = mgl32.Vec2{tX + tWidth, tY}

	// Bottom right
	vertices[5] = mgl32.Vec2{sX + sWidth, sY}
	texs[5] = mgl32.Vec2{tX + tWidth, tY + tHeight}

	return vertices, texs, nil
}

func Square(x, y, width, tX, tY, tWidth int, texture *gl.Texture, window *gl.Window) ([]mgl32.Vec2, []mgl32.Vec2, error) {
	return Rectangle(x, y, width, width, tX, tY, tWidth, tWidth, texture, window)
}
