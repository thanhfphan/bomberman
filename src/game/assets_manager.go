package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type AssetKeeper struct {
	TileTextures map[TileType]*ebiten.Image
}

func NewAssetKeeper() *AssetKeeper {
	return &AssetKeeper{
		TileTextures: make(map[TileType]*ebiten.Image),
	}
}

func (am *AssetKeeper) Load() {
	am.TileTextures[TileEmpty] = grassSprite.GetFrame(0, 0)
	// am.TileTextures[TileGrass] = grassSprite.GetFrame(0, 0)
}
