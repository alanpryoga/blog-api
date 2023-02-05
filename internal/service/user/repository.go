package user

import (
	"context"

	"github.com/alanpryoga/blog-api/internal/entity"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=user

type dbRepoProvider interface {
	FindUserByID(ctx context.Context, id int32) (entity.User, error)
}
