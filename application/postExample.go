package application

import (
	"template_ddd_service/domain/client"

	"github.com/Autoshi-k/context"
	"github.com/Autoshi-k/servicereply"
)

func (app Application) PostExample(ctx context.Context, ID int, name string) (client.PostExampleResponse, servicereply.ServiceReplyI) {
	var response client.PostExampleResponse
	err := app.repo.DoRepositoryThing(ctx)
	if err != nil {
		return response, servicereply.NewInternalServiceError(err)
	}

	response.ID = 1
	return response, nil
}
