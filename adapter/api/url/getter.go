package url

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reptile/domain/model"
	"strconv"
)

const (
	Header  = "header"
	Status  = "status"
	Minutia = "minutia"
)

var instance *PokemonGetter

func Instance() *PokemonGetter {
	if instance == nil {
		instance = &PokemonGetter{}
	}

	return instance
}

type PokemonGetter struct {
}

func (getter PokemonGetter) pokemonDetailHtml(number int) map[string]*goquery.Selection {
	res, err := getter.send(number)
	defer res.Body.Close()
	return getter.selections(err, res)
}

func (getter PokemonGetter) selections(err error, res *http.Response) map[string]*goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	detailSelection := doc.Find(".detail-panel")
	headerSelection := detailSelection.Find(".detail-panel-header")
	statusRowSelection := detailSelection.Find(".detail-stats-row")
	minutiaSelection := detailSelection.Find(".monster-minutia")
	return map[string]*goquery.Selection{"header": headerSelection,
		"status":  statusRowSelection,
		"minutia": minutiaSelection}
}

func (getter PokemonGetter) send(number int) (*http.Response, error) {
	url := fmt.Sprintf("%s%d", "https://pokedex.org/#/pokemon/", number)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return res, err
}
func (getter PokemonGetter) Pokemons() []model.Pokemon {
	var pokemons []model.Pokemon
	for i := 1; i < 650; i++ {
		selectionMap := getter.pokemonDetailHtml(i)
		pokemons[i-1] = getter.generatorHtmlToPokemon(selectionMap, i)
	}
	return pokemons
}

func (getter PokemonGetter) generatorHtmlToPokemon(selectionMap map[string]*goquery.Selection, id int) model.Pokemon {
	headerSelection := selectionMap[Header]
	statusSelection := selectionMap[Status]
	minutiaSelection := selectionMap[Minutia]

	name := headerSelection.Nodes[0].LastChild.Data
	minutiaMap := getter.minutiaMap(minutiaSelection)
	statusMap := getter.statusMap(statusSelection)

	hp, _ := strconv.Atoi(statusMap["HP"])
	attack, _ := strconv.Atoi(statusMap["Attack"])
	defense, _ := strconv.Atoi(statusMap["Defense"])
	speed, _ := strconv.Atoi(statusMap["speed"])
	spAtk, _ := strconv.Atoi(statusMap["Sp Atk"])
	spDef, _ := strconv.Atoi(statusMap["Sp Def"])

	pokemon := model.Pokemon{
		Id:    id,
		Basic: model.Basic{Height: minutiaMap["Height:"], Weight: minutiaMap["Weight:"]},
		Power: model.Power{Name: name, Hp: hp, Attack: attack, Defense: defense, Speed: speed, SpAtk: spAtk, SpDef: spDef},
	}
	return pokemon
}

func (getter PokemonGetter) statusMap(statusSelection *goquery.Selection) map[string]string {
	statusMap := make(map[string]string)
	statusSelection.Each(func(i int, selection *goquery.Selection) {
		first := selection.Nodes[0].FirstChild
		current := first
		for current != nil {
			statusMap[current.FirstChild.Data] = current.NextSibling.LastChild.FirstChild.Data
			current = current.NextSibling.NextSibling
		}
	})
	return statusMap
}

func (getter PokemonGetter) minutiaMap(minutiaSelection *goquery.Selection) map[string]string {
	minutiaMap := make(map[string]string)
	minutiaSelection.Each(func(i int, selection *goquery.Selection) {
		first := selection.Nodes[0].FirstChild
		current := first
		for current != nil {
			if current.NextSibling.FirstChild != nil {
				minutiaMap[current.FirstChild.Data] = current.NextSibling.FirstChild.Data
			}
			current = current.NextSibling.NextSibling
		}
	})
	return minutiaMap
}
