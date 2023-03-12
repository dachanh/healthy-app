package api

import (
	dbmodels "healthyapp/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type mealHistoryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type mealHistorySessionRequest struct {
	ID      int32 `uri:"id" binding:"required,min=1"`
	Session int16 `uri:"session" binding:"required,min=1"`
}

func (s *Server) listMealHistorybySession(ctx *gin.Context) {
	var req mealHistorySessionRequest
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	got, err := s.store.GetMealHistoryBySession(ctx,
		dbmodels.GetMealHistoryBySessionParams{
			UserID:  req.ID,
			Session: req.Session,
			Limit:   int32(intLimit),
			Offset:  int32(offset),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, got)
}

func (s *Server) listMealHistory(ctx *gin.Context) {
	var req mealHistoryRequest
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	got, err := s.store.GetMealHistory(ctx,
		dbmodels.GetMealHistoryParams{
			UserID: req.ID,
			Limit:  int32(intLimit),
			Offset: int32(offset),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, got)
}
