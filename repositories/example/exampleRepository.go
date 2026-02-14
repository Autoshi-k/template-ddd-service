package example

import (
	"template_ddd_service/domain"
	"template_ddd_service/domain/repositories"
	"template_ddd_service/persistence/exampleUser"

	"github.com/Autoshi-k/context"
	"github.com/Autoshi-k/dbtx"
)

type repository struct {
	db dbtx.ConnectionI
}

func NewExampleRepository(db dbtx.ConnectionI) repositories.ExampleRepository {
	return &repository{db: db}
}

func (r repository) DoRepositoryThing(ctx context.Context) error {
	var res exampleUser.SelectEntry
	// dont really like that, need to improve the whole dbtx package
	err := r.db.SelectOne(ctx, &res, dbtx.ToArguments[exampleUser.SelectEntryColumnName](exampleUser.SelectEntryCdn{exampleUser.ID: 1}))
	return err
}

func (r repository) DoRepositoryThingSelectMany(ctx context.Context) error {
	var res exampleUser.SelectEntries
	err := r.db.SelectMany(ctx, &res, dbtx.ToArguments[exampleUser.SelectEntryColumnName](exampleUser.SelectEntryCdn{exampleUser.ID: 1}))
	return err
}

func (r repository) DoRepositoryThingInsertOne(ctx context.Context, user domain.User) error {
	_, err := r.db.InsertOne(ctx, exampleUser.ToInsertEntry(user))
	return err
}

func (r repository) DoRepositoryThingInsertMany(ctx context.Context, users domain.Users) error {
	_, err := r.db.InsertMany(ctx, exampleUser.ToInsertEntries(users))
	return err
}
