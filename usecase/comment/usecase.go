package comment

import (
	"github.com/phamtrung99/community-service/repository"
	"github.com/phamtrung99/community-service/repository/comment"
	userrepository "github.com/phamtrung99/user-service/repository/user"
)

type Usecase struct {
	cmtRepo  comment.Repository
	userRepo userrepository.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		cmtRepo:  repo.Comment,
		userRepo: repo.User,
	}
}
