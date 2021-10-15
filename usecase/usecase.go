package usecase

import (
	"github.com/phamtrung99/community-service/repository"
	"github.com/phamtrung99/community-service/usecase/comment"
)

type UseCase struct {
	Comment comment.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Comment: comment.New(repo),
	}
}
