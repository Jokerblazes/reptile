package application

import (
	"fmt"
	"reptile/adapter/api/selenium"
	"reptile/adapter/db"
	"time"
)

func Reptile() {
	start := time.Now()

	getter := selenium.PokemonGetter{}
	pokemons := getter.Pokemons()

	var saver = db.Repository{}
	saver.Save(pokemons)
	end := time.Now()

	fmt.Println(end.Sub(start))
}
