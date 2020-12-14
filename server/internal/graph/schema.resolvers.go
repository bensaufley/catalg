package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bensaufley/catalg/server/internal/graph/generated"
	"github.com/bensaufley/catalg/server/internal/models"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Find(&users)

	return users, nil
}

func (r *queryResolver) FindUser(ctx context.Context, query string) (*models.User, error) {
	var user *models.User
	tmpl := "%"+query+"%"
	r.DB.Where("username LIKE ? OR email LIKE ?", tmpl).Find(&user)

	return user, nil
}

func (r *queryResolver) User(ctx context.Context, uuid string) (*models.User, error) {
	var user *models.User
	r.DB.Where("uuid = ?", uuid).Find(&user)
	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
