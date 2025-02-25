package powerschooltrackerimport

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"time"
)

type MissingAssignment struct {
	Course     string
	DueDate    string
	Assignment string
}

func Track(username string, password string) {
	var service *selenium.Service
	var err error
	service, err = selenium.NewChromeDriverService("../chromedriver/chromedriver.exe", 4444)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Stop()

	var driver selenium.WebDriver
	driver, err = selenium.NewRemote(selenium.Capabilities{"browserName": "chrome"}, "")
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Quit()

	err = driver.Get("https://scstn.powerschool.com/public/")
	if err != nil {
		log.Fatal(err)
	}

	var formUsername selenium.WebElement
	formUsername, err = driver.FindElement(selenium.ByID, "fieldAccount")
	if err != nil {
		log.Fatal(err)
	}

	var formPassword selenium.WebElement
	formPassword, err = driver.FindElement(selenium.ByID, "fieldPassword")
	if err != nil {
		log.Fatal(err)
	}

	err = formUsername.SendKeys(username)
	if err != nil {
		log.Fatal(err)
	}

	err = formPassword.SendKeys(password)
	if err != nil {
		log.Fatal(err)
	}

	var loginButton selenium.WebElement
	loginButton, err = driver.FindElement(selenium.ByID, "btn-enter-sign-in")
	if err != nil {
		log.Fatal(err)
	}

	err = loginButton.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	var missAssignments selenium.WebElement
	missAssignments, err = driver.FindElement(selenium.ByID, "btn-MissAsmt")
	if err != nil {
		log.Fatal(err)
	}

	err = missAssignments.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	var q3Link selenium.WebElement
	q3Link, err = driver.FindElement(selenium.ByXPATH, "//a[contains(@href, 'missingasmts.html?frn=&trm=Q3')]")
	if err != nil {
		log.Fatal(err)
	}

	err = q3Link.Click()
	if err != nil {
		log.Fatal(err)
	}

	var rows []selenium.WebElement
	rows, err = driver.FindElements(selenium.ByXPATH, "//table[@class='grid']/tbody/tr")
	if err != nil {
		log.Fatal(err)
	}

	var missingAssignments []MissingAssignment

	for _, row := range rows {
		var cols []selenium.WebElement
		cols, err = row.FindElements(selenium.ByTagName, "td")
		if err != nil || len(cols) < 3 {
			log.Printf("Skipping row due to missing columns: %v", err)
			continue
		}
		var course string
		course, _ = cols[0].Text()
		var dueDate string
		dueDate, _ = cols[1].Text()
		var assignment string
		assignment, _ = cols[2].Text()
		missingAssignments = append(missingAssignments, MissingAssignment{
			Course:     course,
			DueDate:    dueDate,
			Assignment: assignment,
		})
	}

	var file *os.File
	file, err = os.Create("../missing.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, m := range missingAssignments {
		_, err = file.WriteString(fmt.Sprintf("Course: %s | Due Date: %s | Assignment: %s\n", m.Course, m.DueDate, m.Assignment))
		if err != nil {
			log.Fatal(err)
		}
	}
}
