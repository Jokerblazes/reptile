package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func StartService() (*selenium.Service, error) {
	opts := []selenium.ServiceOption{}

	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
	}
	return service, err
}

func WebView() (selenium.WebDriver, error) {

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
	return webView, err
}
