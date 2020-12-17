package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/bensaufley/catalg/server/internal/graph/generated"
	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user *models.CreateUserParams) (*models.User, error) {
	return models.CreateUser(
		ctx,
		r.DB,
		models.User{Username: user.Username, Email: user.Email},
		user.Password,
	)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user *models.UpdateUserParams) (*models.User, error) {
	return models.UpdateUser(
		ctx,
		r.DB,
		user.UUID,
		user.Password,
		models.UserUpdateParams{Username: user.Username, Email: user.Email, Password: user.NewPassword},
	)
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Find(&users)

	return users, nil
}

func (r *queryResolver) FindUsers(ctx context.Context, query string) ([]*models.User, error) {
	if query == "" {
		return []*models.User{}, errors.New("query is required")
	}
	var users []*models.User
	tmpl := "%" + query + "%"
	log.WithField("tmpl", tmpl).Debug("about to query")
	if tx := r.DB.WithContext(ctx).Where("username LIKE ? OR email LIKE ?", tmpl, tmpl).Find(&users); tx.Error != nil {
		log.WithError(tx.Error).WithField("query", query).Warn("error looking up users by query")
		return nil, errors.New("could not find users for that query")
	} else if tx.RowsAffected < 1 {
		return []*models.User{}, nil
	}

	return users, nil
}

func (r *queryResolver) GetUser(ctx context.Context, uuid string) (*models.User, error) {
	var user models.User
	if tx := r.DB.WithContext(ctx).Where("uuid = ?", uuid).Find(&user); tx.Error != nil {
		log.WithError(tx.Error).WithField("uuid", uuid).Warn("error looking up user by UUID")
		return nil, errors.New("uuid not valid")
	} else if tx.RowsAffected < 1 {
		return nil, nil
	}
	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
