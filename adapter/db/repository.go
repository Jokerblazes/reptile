package db

import "reptile/domain/model"

type Repository struct {
}

func (repository Repository) Pokemons() []model.Pokemon {

	return nil
}

func (repository Repository) Save(pokemons []model.Pokemon) (int, []error) {
	db := Db()
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO `pokemon` (`id`, `name`, `hp`, `attack`, `defense`, `speed`, `sp_atk`, `sp_def`, `height`, `weight`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()
	failedNum := 0
	var errs []error
	for _, pokemon := range pokemons {
		_, err := stmtIns.Exec(pokemon.Id, pokemon.Name, pokemon.Hp, pokemon.Attack, pokemon.Defense, pokemon.Speed, pokemon.SpAtk, pokemon.SpDef, pokemon.Height, pokemon.Weight)
		if err != nil {
			errs = append(errs, err)
			failedNum++
		}
	}
	return len(pokemons) - failedNum, errs
}
