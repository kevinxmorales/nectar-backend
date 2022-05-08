package main

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/kevinmorales/nectar-rest-api/internal/db"
	"gitlab.com/kevinmorales/nectar-rest-api/internal/plant"
	transportHttp "gitlab.com/kevinmorales/nectar-rest-api/internal/transport/http"
	"gitlab.com/kevinmorales/nectar-rest-api/internal/user"
)

func Run() error {
	log.Info("starting up application")
	database, err := db.NewDatabase()
	if err != nil {
		log.Error("FAILED to connect to the database")
		return err
	}
	if err := database.MigrateDB(); err != nil {
		log.Error("FAILED to migrate database")
		return err
	}

	plantService := plant.NewService(database)
	userService := user.NewService(database)
	httpHandler := transportHttp.NewHandler(plantService, userService)

	log.Info("service has successfully started :)")
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	log.Info("Nectar REST API")
	if err := Run(); err != nil {
		log.Error(err)
	}
}
