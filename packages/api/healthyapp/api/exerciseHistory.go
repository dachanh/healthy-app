package api

import (
	dbmodels "healthyapp/db/sqlc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type exceriseHistoryByDateRequest struct {
	ID    int32     `uri:"id" binding:"required,min=1"`
	Start time.Time `uri:"start" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	End   time.Time `uri:"end" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
}

type exceriseHistoryOfUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getExercseHistoryByDate(ctx *gin.Context) {
	var req exceriseHistoryByDateRequest

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	got, err := s.store.GetExercseHistoryByDate(ctx,
		dbmodels.GetExercseHistoryByDateParams{
			UserID: req.ID,
			Date:   req.Start,
			Date_2: req.End,
			Limit:  int32(intLimit),
			Offset: int32(offset),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, got)
	return
}

func (s *Server) getExercseHistoryOfUser(ctx *gin.Context) {
	var req exceriseHistoryOfUserRequest
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	got, err := s.store.GetExercseHistoryOfUser(ctx,
		dbmodels.GetExercseHistoryOfUserParams{
			UserID: req.ID,
			Limit:  int32(intLimit),
			Offset: int32(offset),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, got)
	return
}
