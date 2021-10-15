package comment

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/community-service/usecase/comment"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
)

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	commentID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	
	req := comment.DeleteCmtRequest{
		ID: commentID,
	}

	err := r.commentUseCase.Delete(ctx, &req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, nil)
}
