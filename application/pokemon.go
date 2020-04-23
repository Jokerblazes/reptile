package application

import (
	"reptile/adapter/api/selenium"
	"reptile/adapter/db"
)

func Reptile() {
	getter := selenium.PokemonGetter{}
	pokemons := getter.Pokemons()

	var saver = db.Repository{}
	saver.Save(pokemons)
}
