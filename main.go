package main

import (
	"fmt"
	"reptile/adapter/api"
)

func main() {
	getter := api.UrlPokemonGetter{}
	pokemons := getter.Pokemons()
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
}
