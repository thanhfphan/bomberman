package spritesheet

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"thanhfphan.com/bomberman/src/engine/math"
)

type SpriteSheet struct {
	Texture    *ebiten.Image
	Width      float64
	Height     float64
	CellWidth  float64
	CellHeight float64
}

func NewSpriteSheet(data []byte, cellWidth, cellHeight float64) (*SpriteSheet, error) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return &SpriteSheet{
		Texture:    img,
		Width:      float64(img.Bounds().Dx()),
		Height:     float64(img.Bounds().Dy()),
		CellWidth:  cellWidth,
		CellHeight: cellHeight,
	}, nil
}

func (s *SpriteSheet) GetFrame(row, column float64) *ebiten.Image {
	x := int(column * s.CellWidth)
	y := int(row * s.CellHeight)
	w := int(s.CellWidth)
	h := int(s.CellHeight)

	return s.Texture.SubImage(image.Rect(x, y, x+w, y+h)).(*ebiten.Image)
}

func (s *SpriteSheet) DrawFrame(screen *ebiten.Image, row, column float64, position math.Vec2, isFlipped bool) {
	frame := s.GetFrame(row, column)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-s.CellWidth/2, -s.CellHeight/2) // Center the image
	if isFlipped {
		op.GeoM.Scale(-1, 1) // flip horizontally
	}
	op.GeoM.Translate(position.X, position.Y)

	screen.DrawImage(frame, op)
}
