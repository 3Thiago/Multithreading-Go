package main

import (
	"fmt"
	"net/http"

	"github.com/3Thiago/Multithreading-Go/configs"
	"github.com/3Thiago/Multithreading-Go/internal/infra/webserver/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	Configs, err := configs.LoadEnvironmentVariables(".")
	if err != nil {
		panic(err)
	}
	logrus.Println(fmt.Sprintf("RUNNING SERVER APP IN PORT: %s", Configs.WEBSERVEPORT))
	handler := handler.NewApiHandler()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/consulta/{cep}", handler.GetCepHandler)

	http.ListenAndServe(":8000", r)
}
