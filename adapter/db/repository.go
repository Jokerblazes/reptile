package db

import (
	"fmt"
	"reptile/domain/model"
)

type Repository struct {
}

func (repository Repository) Pokemons() []model.Pokemon {

	return nil
}

func (repository *Repository) Save(pokemonChan chan model.Pokemon) (chan bool, error) {
	done := make(chan bool)

	go func() {
		db := Db()
		defer db.Close()
		stmtIns, _ := db.Prepare("INSERT INTO `pokemon` (`id`, `name`, `hp`, `attack`, `defense`, `speed`, `sp_atk`, `sp_def`, `height`, `weight`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		defer stmtIns.Close()
		for pokemon := range pokemonChan {
			fmt.Println(pokemon.Name)
			stmtIns.Exec(pokemon.Id, pokemon.Name, pokemon.Hp, pokemon.Attack, pokemon.Defense, pokemon.Speed, pokemon.SpAtk, pokemon.SpDef, pokemon.Height, pokemon.Weight)
		}
		done <- true
	}()
	return done, nil
}
