package comment

import (
	"context"

	"github.com/phamtrung99/community-service/model"
)

// IUsecase .
type IUsecase interface {
	Insert(ctx context.Context, req InsertCmtRequest) (*model.Comment, error)
	Delete(ctx context.Context, req *DeleteCmtRequest) error
	GetList(ctx context.Context, req *GetListRequest) (*model.CommentResult, error)
}
