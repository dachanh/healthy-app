package api

import (
	"database/sql"
	"fmt"
	dbmodels "healthyapp/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type loginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (s *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.Wrap(err, "Failed validate")))
		return
	}

	account, err := s.store.LoginAccount(ctx, dbmodels.LoginAccountParams{
		Email:    req.Email,
		Password: req.Password,
	})
	fmt.Println(req.Email, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err := errors.Wrap(err, "Login not successfully")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SimpleSuccessResponse(account))
}
