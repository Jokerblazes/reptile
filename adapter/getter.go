package adapter

import "reptile/domain/model"

type PokemonGetter interface {
	Pokemons() chan model.Pokemon
}
