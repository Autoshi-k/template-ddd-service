package repositories

import "github.com/Autoshi-k/context"

type ExampleRepository interface {
	DoRepositoryThing(ctx context.Context) error
}
