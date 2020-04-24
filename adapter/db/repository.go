package db

import "reptile/domain/model"

type Repository struct {
}

func (repository Repository) Pokemons() []model.Pokemon {

	return nil
}

func (repository *Repository) Save(pokemonChan chan model.Pokemon) error {
	db := Db()
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO `pokemon` (`id`, `name`, `hp`, `attack`, `defense`, `speed`, `sp_atk`, `sp_def`, `height`, `weight`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	for pokemon := range pokemonChan {
		stmtIns.Exec(pokemon.Id, pokemon.Name, pokemon.Hp, pokemon.Attack, pokemon.Defense, pokemon.Speed, pokemon.SpAtk, pokemon.SpDef, pokemon.Height, pokemon.Weight)
	}
	return nil
}
