package adapter

import "reptile/domain/model"

type UrlPokemonGetter struct {
	url string
}

func (getter UrlPokemonGetter) pokemons() []model.Pokemon {
	pokemon := model.Pokemon{
		BasicProperty: model.BasicProperty{},
		PowerProperty: model.PowerProperty{},
	}

	return []model.Pokemon{pokemon}
	//url := getter.url
	//_,err := http.Get(url)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr,"Error fetch request,url is %s.error is %v",url,err)
	//	os.Exit(-1)
	//}

}
