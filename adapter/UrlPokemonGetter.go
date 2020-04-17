package adapter

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reptile/domain/model"
)

type UrlPokemonGetter struct {
	url string
}

func (getter UrlPokemonGetter) pokemonDetailHtml(number int) map[string]*goquery.Selection {
	url := fmt.Sprintf("%s%d", "https://pokedex.org/#/pokemon/", number)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	detailSelection := doc.Find(".detail-panel")
	headerSelection := detailSelection.Find(".detail-panel-header")
	statusRowSelection := detailSelection.Find(".detail-stats-row")
	minutiaSelection := detailSelection.Find(".monster-minutia")
	return map[string]*goquery.Selection{"header": headerSelection, "status": statusRowSelection, "minutia": minutiaSelection}
}
func (getter UrlPokemonGetter) pokemons() []model.Pokemon {

	selectionMap := getter.pokemonDetailHtml(1)
	return getter.generatorHtmlToPokemon(selectionMap)
}

func (getter UrlPokemonGetter) generatorHtmlToPokemon(selectionMap map[string]*goquery.Selection) []model.Pokemon {
	pokemon := model.Pokemon{
		BasicProperty: model.BasicProperty{},
		PowerProperty: model.PowerProperty{},
	}

	return []model.Pokemon{pokemon}
}
