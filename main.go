package main

import (
	"os"
	"strings"

	"github.com/babamba/learngo/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	//return c.String(http.StatusOK, "Hello, World!")
	term := strings.ToLower(scrapper.Cleanstring(c.FormValue("term")))
	//return c.File("home.html")
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	//scrapper.Scrape("frontend")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
