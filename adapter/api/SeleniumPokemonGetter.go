package api

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"reptile/domain/model"
	"time"
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

func (getter SeleniumPokemonGetter) pokemonDetailHtml(number int) interface{} {
	return nil
}

func (getter SeleniumPokemonGetter) generatorHtmlToPokemon(selectionMap interface{}, number int) model.Pokemon {
	return model.Pokemon{}
}

func pokemonDetailHtml() {
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

	// 調起chrome瀏覽器
	w_b1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	//關閉一個webDriver會對應關閉一個chrome窗口
	//但是不會導致seleniumServer關閉
	defer w_b1.Quit()
	err = w_b1.Get("https://pokedex.org/#/pokemon/1")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}

	// 重新調起chrome瀏覽器
	w_b2, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	defer w_b2.Close()
	//打開一個網頁
	err = w_b2.Get("https://www.toutiao.com/")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	//打開一個網頁
	err = w_b2.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	//w_b就是當前頁面的對象，通過該對象可以操作當前頁面了
	//........
	time.Sleep(5 * time.Minute)
	return
}
