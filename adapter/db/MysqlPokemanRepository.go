package db

import "reptile/domain/model"

type MysqlPokemanRepository struct {
}

func (repository MysqlPokemanRepository) Pokemons() []model.Pokemon {
	return nil
}

func (repository MysqlPokemanRepository) savePokemons([]model.Pokemon) int {
	return 0
}
