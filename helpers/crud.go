package helpers

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"tutorial.sqlc.dev/app/tutorial"
)

func CreateAuthor(ctx context.Context, query *tutorial.Queries, name string, bio string) (*tutorial.Author, error) {
	newAuthor, err := query.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: name,
		Bio:  pgtype.Text{String: bio},
	})
	if err != nil {
		return nil, err
	}
	return &newAuthor, nil
}

func ReadAuthor(ctx context.Context, query *tutorial.Queries, id int64) (*tutorial.Author, error) {
	author, err := query.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func ReadAuthors(ctx context.Context, query *tutorial.Queries) ([]tutorial.Author, error) {
	authors, err := query.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func UpdateAuthorName(ctx context.Context, query *tutorial.Queries, id int64, name string) (*tutorial.Author, error) {
	author, err := query.UpdateAuthor(ctx, tutorial.UpdateAuthorParams{
		ID:   id,
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func UpdateAuthorBio(ctx context.Context, query *tutorial.Queries, id int64, bio string) (*tutorial.Author, error) {
	author, err := query.UpdateAuthor(ctx, tutorial.UpdateAuthorParams{
		ID:  id,
		Bio: pgtype.Text{String: bio},
	})
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func DeleteAuthor(ctx context.Context, query *tutorial.Queries, id int64) error {
	if err := query.DeleteAuthor(ctx, id); err != nil {
		return err
	}
	return nil
}
