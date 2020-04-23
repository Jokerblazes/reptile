package adapter

import "reptile/domain/model"

type PokemonModifier interface {
	savePokemons([]model.Pokemon) (int, []error)
}
