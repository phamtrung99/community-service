package comment

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/community-service/model"
	"github.com/phamtrung99/community-service/usecase/comment"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	moviemodel "github.com/phamtrung99/movie-service/model"
)

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
	)

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	movieID, _ := strconv.ParseInt(c.QueryParam("movie_id"), 10, 64)
	parentID, _ := strconv.ParseInt(c.QueryParam("parent_id"), 10, 64)

	req := comment.GetListRequest{
		Paginator: &moviemodel.Paginator{
			Page:  page,
			Limit: limit,
		},
		Filter: &model.CommentFilter{
			MovieID:  movieID,
			ParentID: parentID,
		},
	}

	res, err := r.commentUseCase.GetList(ctx, &req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
