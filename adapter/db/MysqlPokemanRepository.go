package db

import "reptile/domain/model"

type MysqlPokemanRepository struct {
}

func (repository MysqlPokemanRepository) Pokemons() []model.Pokemon {

	return nil
}

func (repository MysqlPokemanRepository) savePokemons(pokemons []model.Pokemon) int {
	db := Db()
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO `pokemon` (`id`, `name`, `hp`, `attack`, `defense`, `speed`, `sp_atk`, `sp_def`, `height`, `weight`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	failedNum := 0
	for _, pokemon := range pokemons {
		_, err := stmtIns.Exec(pokemon.Id, pokemon.Name, pokemon.Hp, pokemon.Attack, pokemon.Speed, pokemon.SpAtk, pokemon.SpDef, pokemon.Height, pokemon.Weight)
		if err != nil {
			failedNum++
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return len(pokemons) - failedNum
}
