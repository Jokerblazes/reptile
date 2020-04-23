package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"reptile/domain/model"
	"strconv"
	"strings"
)

const (
	seleniumPath = `./chromedriver`
	port         = 9515
)

type PokemonGetter struct {
}

func (getter *PokemonGetter) Pokemons() []model.Pokemon {
	var pokemons []model.Pokemon
	service, _ := getter.startService()
	defer service.Stop()
	webView := *getter.getWebView()
	defer webView.Close()

	monsterButtons, _ := webView.FindElements(selenium.ByClassName, "monster-sprite")
	for i, monsterButton := range monsterButtons {
		name, statsMap, minutiaMap := getter.pokemonDetailHtml(monsterButton, webView)
		pokemons = append(pokemons, getter.generatorHtmlToPokemon(name, statsMap, minutiaMap, i+1))
	}
	return pokemons
}

func (getter *PokemonGetter) getWebView() *selenium.WebDriver {

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
	//打開一個網頁
	err = webView.Get("https://pokedex.org/#/")
	if err != nil {
		fmt.Println("get page faild", err.Error())
	}
	return &webView
}

func (getter *PokemonGetter) startService() (*selenium.Service, error) {

	opts := []selenium.ServiceOption{}

	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
	}
	return service, err
}
func (getter *PokemonGetter) pokemonDetailHtml(monsterButton selenium.WebElement, view selenium.WebDriver) (name string, statsMap map[string]string, minutiaMap map[string]string) {
	monsterButton.Click()
	view.Wait(func(wd selenium.WebDriver) (bool, error) {
		value, err := view.FindElement(selenium.ByClassName, "detail-panel-header")
		b, err2, done := getter.resourceReady(err, []selenium.WebElement{value})
		if done {
			return b, err2
		}
		values, err := view.FindElements(selenium.ByCSSSelector, ".detail-stats-row span")
		b2, err3, done2 := getter.resourceReady(err, values)
		if done2 {
			return b2, err3
		}
		values, err = view.FindElements(selenium.ByCSSSelector, ".monster-minutia span")
		b2, err3, done2 = getter.resourceReady(err, values)
		if done2 {
			return b2, err3
		}
		values, err = view.FindElements(selenium.ByCSSSelector, ".monster-minutia strong")
		b2, err3, done2 = getter.resourceReady(err, values)
		if done2 {
			return b2, err3
		}
		values, err = view.FindElements(selenium.ByClassName, "stat-bar-fg")
		b2, err3, done2 = getter.resourceReady(err, values)
		if done2 {
			return b2, err3
		}
		return true, nil
	})

	defer func() {
		backButton, _ := view.FindElement(selenium.ByClassName, "back-button")
		backButton.Click()
	}()
	minutiaMap = make(map[string]string)
	statsMap = make(map[string]string)
	headerElement, _ := view.FindElement(selenium.ByClassName, "detail-panel-header")
	statsElements, _ := view.FindElements(selenium.ByClassName, "detail-stats-row")
	minutiaElements, _ := view.FindElements(selenium.ByClassName, "monster-minutia")
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

func (getter *PokemonGetter) resourceReady(err error, values []selenium.WebElement) (bool, error, bool) {
	if err != nil {
		return false, err, true
	}
	for _, value := range values {
		text, err := value.Text()
		if err != nil || strings.EqualFold(text, "") {
			return false, err, true
		}
	}
	return false, nil, false
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
