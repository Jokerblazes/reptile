package adapter

import "reptile/domain/model"

type PokemonGetter interface {
	pokemons(url string) []model.Pokemon
}
