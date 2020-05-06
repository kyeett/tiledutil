package tiledutil

import (
	"github.com/peterhellberg/gfx"
	"image"
	"log"
	"path/filepath"
)

func (mp *MapWrapper) Image(f ...Filter) (image.Image, error) {
	var filter Filter = NoFilter
	if len(f) > 0 {
		filter = f[0]
	}
	path := filepath.Join(mp.directory, mp.Tilesets[0].Image.Source)

	src, err := gfx.DecodeImageBytes(mp.getAssets(path))
	if err != nil {
		return nil, err
	}

	dst := gfx.NewImage(mp.Width*mp.TileWidth, mp.Height*mp.TileHeight)

	// Load images
	tileSize := gfx.IR(0, 0, mp.TileWidth, mp.TileHeight)
	for _, layer := range mp.Layers {
		for i, tile := range layer.Tiles {
			if !tile.IsNil() {
				tTile := mp.tilesetTileLookup[tile.ID]
				if !filter(layer, tile, tTile.Type) {
					continue
				}
				x := i % mp.Width
				y := i / mp.Width

				tx := int((tile.ID) % uint32(tile.Tileset.Columns))
				ty := int((tile.ID) / uint32(tile.Tileset.Columns))

				altRect := tileSize.Add(image.Pt(mp.TileHeight*x, mp.TileWidth*y))
				altPt := image.Pt(mp.TileWidth*tx, mp.TileHeight*ty)

				gfx.DrawSrc(dst, altRect, src, altPt)
			}
		}
	}
	return dst, nil
}

func (mp *MapWrapper) MustImage(f ...Filter) image.Image {
	dst, err := mp.Image(f...)
	if err != nil {
		log.Fatal(err)
	}
	return dst
}
