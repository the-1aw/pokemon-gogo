package pokeapi

import (
	"time"

	"github.com/the-1aw/pokemon-gogo/internal/pokecache"
)

var PokeCache = pokecache.NewCache(2 * time.Minute)

