package adapter

import "reptile/domain/model"

type PokemonModifier interface {
	savePokemons(pokemons chan model.Pokemon) error
}
