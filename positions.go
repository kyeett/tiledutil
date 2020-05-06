package tiledutil

import (
	"github.com/lafriks/go-tiled"
)

type Position struct {
	X, Y int64
}

type Filter func(l *tiled.Layer, tile *tiled.LayerTile, typ string) bool

var NoFilter = func(_ *tiled.Layer, _ *tiled.LayerTile, _ string) bool { return true }

func (mp *MapWrapper) Positions(f ...Filter) []Position {
	var filter Filter = NoFilter
	if len(f) > 0 {
		filter = f[0]
	}

	var positions []Position
	for _, layer := range mp.Layers {
		for i, tile := range layer.Tiles {
			if !tile.IsNil() {
				tTile := mp.tilesetTileLookup[tile.ID]
				if !filter(layer, tile, tTile.Type) {
					continue
				}

				x := i % mp.Width
				y := i / mp.Width
				positions = append(positions, Position{
					X: int64(x),
					Y: int64(y),
				})
			}
		}
	}
	return positions
}

type TileInfo struct {
	*tiled.LayerTile
	*tiled.TilesetTile
	X int
	Y int
}
