package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
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

	body, _ := io.ReadAll(resp.Body)

	logger.Info(string(body), "user", os.Getenv("USER"))
}
