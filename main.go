package main

import (
	"fmt"
	"reptile/adapter"
)

func main() {
	getter := adapter.UrlPokemonGetter{}
	pokemons := getter.Pokemons()
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
}
