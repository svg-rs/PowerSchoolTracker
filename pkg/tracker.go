package powerschooltracker

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

func Track(username int, password string) {
	service, err := selenium.NewChromeDriverService("../chromedriver/chromedriver.exe", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://scstn.powerschool.com/guardian/classassignments.html")
	if err != nil {
		log.Fatal("Error:", err)
	}

	html, err := driver.PageSource()
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(html)
}
