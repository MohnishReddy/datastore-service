package pkg

import (
	"datastore-service/constants"
	"datastore-service/models"
)

type App struct {
	Config *models.Config
}

var app *App

func InitApp(env constants.ServiceMode) error {
	config, err := ReadConfig(env)
	if err != nil {
		return err
	}

	app = &App{
		Config: config,
	}
	return nil
}

func GetConfig() *models.Config {
	return app.Config
}
