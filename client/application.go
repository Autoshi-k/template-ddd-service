package client

import (
	"log"
	"net/http"
	"template_ddd_service/domain/application"
	"template_ddd_service/domain/client"

	"github.com/Autoshi-k/clientcore"
	"github.com/Autoshi-k/context"
	"github.com/Autoshi-k/servicereply"
)

type applicationClient struct {
	app application.ApplicationI
}

func RunApplicationClient(app application.ApplicationI) {
	c := applicationClient{app: app}

	http.HandleFunc("/postExample", clientcore.POSTRequest[client.PostExampleRequest, client.PostExampleResponse](c.PostExample))
	http.HandleFunc("/getExample", clientcore.GETRequest[client.GetExampleRequest, client.GetExampleResponse](c.GetExample))

	log.Println("Serving on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (c *applicationClient) PostExample(ctx context.Context, request client.PostExampleRequest) (response client.PostExampleResponse, err servicereply.ServiceReplyI) {
	return c.app.PostExample(ctx, request.ID, request.Name)
}

func (c *applicationClient) GetExample(ctx context.Context, request client.GetExampleRequest) (response client.GetExampleResponse, err servicereply.ServiceReplyI) {
	return c.app.GetExample(ctx, request.ID)
}
