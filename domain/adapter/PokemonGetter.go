package adapter

import "reptile/domain/model"

type PokemonGetter interface {
	pokemons() []model.Pokemon
}
