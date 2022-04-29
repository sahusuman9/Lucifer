package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

// var quotes []quote = make([]quote, 0)

var quotes []quote = []quote{
	{"Learn to Lead", "Sai Vidya Campus"},
	{"You can totally do this", "Unknown"},
	{"You are enough", "Unknown"},
	{"You are great", "Unknown"},
	{"You have the potential to be anything and do anything", "Unknown"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":" + port)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}
func getQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}
