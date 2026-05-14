package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	//api.RegisterRoutes(mux, logger)

	handler :-
}