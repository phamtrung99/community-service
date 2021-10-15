package comment

import (
	"context"

	"github.com/phamtrung99/community-service/model"
	moviemodel "github.com/phamtrung99/movie-service/model"
)

// Repository .
type Repository interface {
	Find(
		ctx context.Context,
		conditions []moviemodel.Condition,
		paginator *moviemodel.Paginator,
		orders []string,
	) (*model.CommentResult, error)
	Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id int64) error
	DeleteSubCmt(ctx context.Context, parentID int64) error
	GetById(ctx context.Context, id int64) (*model.Comment, error)
}
