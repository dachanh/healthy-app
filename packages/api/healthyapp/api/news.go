package api

import (
	dbmodels "healthyapp/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) listNews(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	got, err := s.store.ListNews(ctx, dbmodels.ListNewsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, SimpleSuccessResponse(got))
}
