package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"reptile/domain/model"
)

type SeleniumPokemonGetter struct {
}

func (getter SeleniumPokemonGetter) Pokemons() []model.Pokemon {
	pokemons := make([]model.Pokemon, 649)
	for i := 1; i < 650; i++ {
		selectionMap := getter.pokemonDetailHtml(i)
		pokemons[i-1] = getter.generatorHtmlToPokemon(selectionMap, i)
	}
	return pokemons
}

func (getter SeleniumPokemonGetter) pokemonDetailHtml(number int) map[string]interface{} {
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
	minutiaMap := make(map[string]string)
	statsMap := make(map[string]string)

	headerElement, _ := webView.FindElement(selenium.ByClassName, "detail-panel-header")
	statsElements, _ := webView.FindElements(selenium.ByClassName, "detail-stats-row")
	minutiaElements, _ := webView.FindElements(selenium.ByClassName, "monster-minutia")
	name, _ := headerElement.Text()
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
	return map[string]interface{}{
		"header":  name,
		"status":  statsMap,
		"minutia": minutiaMap,
	}
}

func (getter SeleniumPokemonGetter) generatorHtmlToPokemon(pokemonMap map[string]interface{}, number int) model.Pokemon {
	return model.Pokemon{}
}

func PokemonDetailHtml() {
	const (
		seleniumPath = `./chromedriver`
		port         = 9515
	)

	opts := []selenium.ServiceOption{}

	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
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
		return
	}
	defer webView.Close()
	//打開一個網頁
	err = webView.Get("https://pokedex.org/#/pokemon/1")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	minutiaMap := make(map[string]string)
	statsMap := make(map[string]string)

	headerElement, _ := webView.FindElement(selenium.ByClassName, "detail-panel-header")
	statsElements, _ := webView.FindElements(selenium.ByClassName, "detail-stats-row")
	minutiaElements, _ := webView.FindElements(selenium.ByClassName, "monster-minutia")
	name, _ := headerElement.Text()
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
	fmt.Println(name)
}
