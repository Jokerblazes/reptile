package adapter

import "reptile/domain/model"

type PokemonModifier interface {
	savePokemons(pokemonChan chan model.Pokemon) (chan bool, error)
}
