package main

import (
	"salon/internal/config"
	"salon/internal/db"
	"salon/internal/handler"
	"salon/internal/logger"
	"salon/internal/repo"
	"salon/internal/router"
	"salon/internal/service"
)

func main() {
	c := config.New()

	l := logger.New(c.Env)

	d := db.New(c.SalonDsn)

	r := repo.New(d)

	s := service.New(r)

	h := handler.New(s)

	server := router.New(c.Port, h)

	// init cancellation ctx

	server.ListenAndServe()
}
