package main

import (
	"template_ddd_service/application"
	"template_ddd_service/client"
	"template_ddd_service/repositories/example"
)

func main() {
	app := application.NewApplication(
		example.NewExampleRepository(),
	)
	client.RunApplicationClient(app)
}
