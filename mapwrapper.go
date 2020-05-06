package tiledutil

import (
	"bytes"
	"github.com/lafriks/go-tiled"
	"io/ioutil"
	"log"
	"path/filepath"
)

type MapWrapper struct {
	*tiled.Map
	directory string
	tilesetTileLookup map[uint32]*tiled.TilesetTile

	getAssets func(path string) []byte
}

func MustFromFile(path string) *MapWrapper {
	dir := filepath.Dir(path)
	mp, err := tiled.LoadFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	getAssets := func(path string) []byte {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		return b
	}

	return baseWrapper(mp, dir, getAssets)
}

func MustFromBytes(path string, getAssets func(path string) []byte) *MapWrapper {
	dir := filepath.Dir(path)
	mp, err := tiled.LoadFromReader(dir, bytes.NewReader(getAssets(path)))
	if err != nil {
		log.Fatal(err)
	}

	return baseWrapper(mp, dir, getAssets)
}

func baseWrapper(mp *tiled.Map, dir string, getAssets func(path string) []byte) *MapWrapper {
	lookup := map[uint32]*tiled.TilesetTile{}
	for _, t := range mp.Tilesets[0].Tiles {
		lookup[t.ID] = t
	}

	w := &MapWrapper{mp, dir, lookup, getAssets}
	return w
}


