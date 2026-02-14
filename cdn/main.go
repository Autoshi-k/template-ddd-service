package main

import (
	"os"
	"template_ddd_service/application"
	"template_ddd_service/client"
	"template_ddd_service/repositories/example"

	"github.com/Autoshi-k/dbtx"
)

func main() {
	app := application.NewApplication(
		example.NewExampleRepository(
			dbtx.NewDatabaseConnection(os.Getenv("PG_CONN")),
		),
	)
	client.RunApplicationClient(app)
}
