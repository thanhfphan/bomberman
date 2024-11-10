package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	TileEmpty TileType = iota
	TileGrass
)

type Tile struct {
	Type TileType
}
type Grid struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

func NewGrid(width, height int) *Grid {
	grid := Grid{
		Width:  width,
		Height: height,
		Tiles:  make([][]Tile, height),
	}
	for i := range grid.Tiles {
		grid.Tiles[i] = make([]Tile, width)
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j] = Tile{
				Type: TileEmpty,
			}
		}
	}
	return &grid
}

func (g *Grid) PlaceTile(x, y int, tile Tile) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		panic(fmt.Sprintf("Invalid tile position: (%d, %d)", x, y))
	}
	g.Tiles[y][x] = tile
}

func (g *Grid) RemoveTile(x, y int) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		panic(fmt.Sprintf("Invalid tile position: (%d, %d)", x, y))
	}
	g.Tiles[y][x] = Tile{
		Type: TileEmpty,
	}
}

func (g *Grid) IsWalkAble(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}

	return g.Tiles[y][x].Type != TileEmpty
}

func (g *Grid) Render(screen *ebiten.Image) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			tile := g.Tiles[y][x]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))

			screen.DrawImage(global.assetKeeper.TileTextures[tile.Type], op)
		}
	}
}
