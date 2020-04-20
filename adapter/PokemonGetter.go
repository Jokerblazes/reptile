package adapter

import "reptile/domain/model"

type PokemonGetter interface {
	Pokemons() []model.Pokemon
}
