package application

import (
	"reptile/adapter"
	"reptile/adapter/api"
	"reptile/adapter/db"
)

func Reptile() {
	var getter adapter.PokemonGetter = api.Instance()
	pokemons := getter.Pokemons()

	var saver = db.Instance()
	saver.SavePokemons(pokemons)
}
