package application

import (
	"template_ddd_service/domain/application"
	"template_ddd_service/domain/repositories"
)

type Application struct {
	repo repositories.ExampleRepository
}

func NewApplication(repo repositories.ExampleRepository) application.ApplicationI {
	app := Application{repo: repo}
	return app
}
