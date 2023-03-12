package api

import (
	dbmodels "healthyapp/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createDiaryRequest struct {
	user_id int32  `json:"user_id" binding:"required"`
	content string `json:"content" binding:"required"`
}

type getDiaryRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *Server) createDiary(ctx *gin.Context) {
	var req createDiaryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := s.store.CreateDiary(ctx, dbmodels.CreateDiaryParams{
		UserID:  req.user_id,
		Content: req.content,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "create diary succcesfully"})
}

func (s *Server) getUserDiary(ctx *gin.Context) {
	var req getDiaryRequest
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	got, err := s.store.GetUserDiary(ctx,
		dbmodels.GetUserDiaryParams{
			UserID: req.ID,
			Limit:  int32(intLimit),
			Offset: int32(offset),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, got)
	return
}
