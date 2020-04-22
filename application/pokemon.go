package application

import (
	"reptile/adapter/api/url"
	"reptile/adapter/db"
)

func Reptile() {
	getter := url.PokemonGetter{}
	pokemons := getter.Pokemons()

	var saver = db.Repository{}
	saver.Save(pokemons)
}
