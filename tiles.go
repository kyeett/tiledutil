package tiledutil

import "fmt"

func (mp *MapWrapper) Tiles(f ...Filter) []TileInfo {
	var filter Filter = NoFilter
	if len(f) > 0 {
		filter = f[0]
	}

	var tileInfo []TileInfo
	for _, layer := range mp.Layers {
		for i, tile := range layer.Tiles {
			if !tile.IsNil() {
				tTile := mp.tilesetTileLookup[tile.ID]
				fmt.Println(tile.ID, tTile)
				if !filter(layer, tile, tTile.Type) {
					continue
				}

				tileInfo = append(tileInfo, TileInfo{
					LayerTile:   tile,
					TilesetTile: tTile,
					X:           i % mp.Width,
					Y:           i / mp.Width,
				})
			}
		}
	}
	return tileInfo
}
