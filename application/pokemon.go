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
	getter.Pokemons()
	pokemonChan := getter.Pokemons()

	var saver = db.Repository{}
	done, _ := saver.Save(pokemonChan)
	fmt.Println(<-done)
	end := time.Now()

	fmt.Println(end.Sub(start))
}
