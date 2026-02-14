package application

import (
	"template_ddd_service/domain/client"

	"github.com/Autoshi-k/context"
	"github.com/Autoshi-k/servicereply"
)

type ApplicationI interface {
	PostExample(ctx context.Context, ID int, name string) (response client.PostExampleResponse, err servicereply.ServiceReplyI)
	GetExample(ctx context.Context, ID int) (response client.GetExampleResponse, err servicereply.ServiceReplyI)
}
