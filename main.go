package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	var catFact CatFact
	// create a logger type
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Hello Terminal From Your Favourite Neighbourhood Backend Developer")

	// create an http client
	client := http.DefaultClient

	resp, err := client.Get("https://catfact.ninja/fact")
	// check for error
	if err != nil {
		logger.Error("response error from getting a cat fact")
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		logger.Error("error occured while decoding json data into catFact struct", "error", err.Error())
		return
	}

	logger.Info(catFact.Fact, "Fact length", catFact.Length)
}
