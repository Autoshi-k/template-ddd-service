package example

import (
	"template_ddd_service/domain/repositories"

	"github.com/Autoshi-k/context"
)

type repository struct {
}

func NewExampleRepository() repositories.ExampleRepository {
	return &repository{}
}

func (r repository) DoRepositoryThing(ctx context.Context) error {
	// todo need to add db support when i figure out the db package
	return nil
}
