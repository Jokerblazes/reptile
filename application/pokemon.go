package application

import (
	"reptile/adapter"
	"reptile/adapter/api/url"
	"reptile/adapter/db"
)

func Reptile() {
	var getter adapter.PokemonGetter = url.Instance()
	pokemons := getter.Pokemons()

	var saver = db.Instance()
	saver.Save(pokemons)
}
