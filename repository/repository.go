package repository

import (
	"context"

	"github.com/phamtrung99/community-service/repository/comment"
	"gorm.io/gorm"
	"github.com/phamtrung99/user-service/repository/user"
)

type Repository struct {
	Comment comment.Repository
	User	user.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		Comment: comment.NewPGRepository(getSQLClient),
		User: user.NewPGRepository(getSQLClient),
	}
}
