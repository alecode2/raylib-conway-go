package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func MakeNineSliceRegions(tex rl.Texture2D, v1, v2, h1, h2 float32) [9]rl.Rectangle {
	texW := float32(tex.Width)
	texH := float32(tex.Height)
	return [9]rl.Rectangle{
		{X: 0, Y: 0, Width: v1, Height: h1},
		{X: v1, Y: 0, Width: v2 - v1, Height: h1},
		{X: v2, Y: 0, Width: texW - v2, Height: h1},

		{X: 0, Y: h1, Width: v1, Height: h2 - h1},
		{X: v1, Y: h1, Width: v2 - v1, Height: h2 - h1},
		{X: v2, Y: h1, Width: texW - v2, Height: h2 - h1},

		{X: 0, Y: h2, Width: v1, Height: texH - h2},
		{X: v1, Y: h2, Width: v2 - v1, Height: texH - h2},
		{X: v2, Y: h2, Width: texW - v2, Height: texH - h2},
	}
}

func DrawNineSlice(
	texture rl.Texture2D,
	src [9]rl.Rectangle,
	dst rl.Rectangle,
	tint rl.Color,
	tileCenter, tileEdges bool,
) {
	left := src[0].Width
	right := src[2].Width
	top := src[0].Height
	bottom := src[6].Height

	midW := dst.Width - left - right
	midH := dst.Height - top - bottom

	// Destination rects
	dsts := [9]rl.Rectangle{
		{X: dst.X, Y: dst.Y, Width: left, Height: top},
		{X: dst.X + left, Y: dst.Y, Width: midW, Height: top},
		{X: dst.X + left + midW, Y: dst.Y, Width: right, Height: top},

		{X: dst.X, Y: dst.Y + top, Width: left, Height: midH},
		{X: dst.X + left, Y: dst.Y + top, Width: midW, Height: midH},
		{X: dst.X + left + midW, Y: dst.Y + top, Width: right, Height: midH},

		{X: dst.X, Y: dst.Y + top + midH, Width: left, Height: bottom},
		{X: dst.X + left, Y: dst.Y + top + midH, Width: midW, Height: bottom},
		{X: dst.X + left + midW, Y: dst.Y + top + midH, Width: right, Height: bottom},
	}

	for i := 0; i < 9; i++ {
		drawSrc := src[i]
		drawDst := dsts[i]

		if i == 4 && tileCenter {
			TileTexture(texture, drawSrc, drawDst, tint)
		} else if (i == 1 || i == 3 || i == 5 || i == 7) && tileEdges {
			TileTexture(texture, drawSrc, drawDst, tint)
		} else {
			rl.DrawTexturePro(texture, drawSrc, drawDst, rl.NewVector2(0, 0), 0, tint)
		}
	}
}

/*
Tiles a texture of size src to fill size dst
*/
func TileTexture(texture rl.Texture2D, src rl.Rectangle, dst rl.Rectangle, tint rl.Color) {
	tilesX := int(dst.Width / src.Width)
	tilesY := int(dst.Height / src.Height)

	remX := dst.Width - float32(tilesX)*src.Width
	remY := dst.Height - float32(tilesY)*src.Height

	for y := 0; y < tilesY; y++ {
		for x := 0; x < tilesX; x++ {
			pos := rl.NewRectangle(
				dst.X+float32(x)*src.Width,
				dst.Y+float32(y)*src.Height,
				src.Width,
				src.Height,
			)
			rl.DrawTexturePro(texture, src, pos, rl.NewVector2(0, 0), 0, tint)
		}
		// Optional: partial tile on X-axis
		if remX > 0 {
			srcClip := rl.NewRectangle(src.X, src.Y, remX, src.Height)
			dstClip := rl.NewRectangle(dst.X+float32(tilesX)*src.Width, dst.Y+float32(y)*src.Height, remX, src.Height)
			rl.DrawTexturePro(texture, srcClip, dstClip, rl.NewVector2(0, 0), 0, tint)
		}
	}

	// Optional: partial tile on Y-axis
	if remY > 0 {
		for x := 0; x < tilesX; x++ {
			srcClip := rl.NewRectangle(src.X, src.Y, src.Width, remY)
			dstClip := rl.NewRectangle(dst.X+float32(x)*src.Width, dst.Y+float32(tilesY)*src.Height, src.Width, remY)
			rl.DrawTexturePro(texture, srcClip, dstClip, rl.NewVector2(0, 0), 0, tint)
		}

		// Bottom-right corner (partial X and Y)
		if remX > 0 {
			srcClip := rl.NewRectangle(src.X, src.Y, remX, remY)
			dstClip := rl.NewRectangle(dst.X+float32(tilesX)*src.Width, dst.Y+float32(tilesY)*src.Height, remX, remY)
			rl.DrawTexturePro(texture, srcClip, dstClip, rl.NewVector2(0, 0), 0, tint)
		}
	}
}
