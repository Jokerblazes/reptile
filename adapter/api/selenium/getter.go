package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"reptile/domain/model"
	"strconv"
)

type PokemonGetter struct {
}

func (getter *PokemonGetter) Pokemons() []model.Pokemon {
	pokemons := make([]model.Pokemon, 649)
	for i := 1; i < 650; i++ {
		name, statsMap, minutiaMap := getter.pokemonDetailHtml(i)
		pokemons[i-1] = getter.generatorHtmlToPokemon(name, statsMap, minutiaMap, i)
	}
	return pokemons
}

func (getter *PokemonGetter) pokemonDetailHtml(number int) (name string, statsMap map[string]string, minutiaMap map[string]string) {
	const (
		seleniumPath = `./chromedriver`
		port         = 9515
	)

	opts := []selenium.ServiceOption{}

	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless",
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模擬user-agent，防反爬
		},
	}
	//以上是設置瀏覽器參數
	caps.AddChrome(chromeCaps)

	// 重新調起chrome瀏覽器
	webView, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
	}
	defer webView.Close()
	//打開一個網頁
	err = webView.Get(fmt.Sprintf("%s%d", "https://pokedex.org/#/pokemon/", number))
	if err != nil {
		fmt.Println("get page faild", err.Error())
	}
	minutiaMap = make(map[string]string)
	statsMap = make(map[string]string)

	headerElement, _ := webView.FindElement(selenium.ByClassName, "detail-panel-header")
	statsElements, _ := webView.FindElements(selenium.ByClassName, "detail-stats-row")
	minutiaElements, _ := webView.FindElements(selenium.ByClassName, "monster-minutia")
	name, _ = headerElement.Text()
	for _, stats := range statsElements {
		keyValueElement, _ := stats.FindElements(selenium.ByTagName, "span")
		key, _ := keyValueElement[0].Text()
		value, _ := keyValueElement[1].Text()
		statsMap[key] = value
	}

	for _, minutia := range minutiaElements {
		keys, _ := minutia.FindElements(selenium.ByTagName, "strong")
		values, _ := minutia.FindElements(selenium.ByTagName, "span")
		for i := 0; i < len(keys); i++ {
			key, _ := keys[i].Text()
			value, _ := values[i].Text()
			minutiaMap[key] = value
		}
	}
	return name, statsMap, minutiaMap
}

func (getter *PokemonGetter) generatorHtmlToPokemon(name string, statsMap map[string]string, minutiaMap map[string]string, number int) model.Pokemon {
	hp, _ := strconv.Atoi(statsMap["HP"])
	attack, _ := strconv.Atoi(statsMap["Attack"])
	defense, _ := strconv.Atoi(statsMap["Defense"])
	speed, _ := strconv.Atoi(statsMap["speed"])
	spAtk, _ := strconv.Atoi(statsMap["Sp Atk"])
	spDef, _ := strconv.Atoi(statsMap["Sp Def"])

	return model.Pokemon{
		Id:    number,
		Basic: model.Basic{Height: minutiaMap["Height:"], Weight: minutiaMap["Weight:"]},
		Power: model.Power{Name: name, Hp: hp, Attack: attack, Defense: defense, Speed: speed, SpAtk: spAtk, SpDef: spDef},
	}
}
