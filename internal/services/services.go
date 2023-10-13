package services

import (
	"net/http"

	errorz "github.com/dungnh3/mfv-codingchallenge/internal/errors"
	"github.com/gin-gonic/gin"
)

type (
	getUserResponse struct {
		ID         int64   `json:"id"`
		Name       string  `json:"name"`
		AccountIDs []int64 `json:"account_ids"`
	}
)

// @Summary get user information
// @Description get user information
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {object} getUserResponse
// @Router /users/{id} [get]
func (s *Server) getUser(ctx *gin.Context) {
	id, err := getIDFromPath(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrIDInvalid, err.Error()))
		return
	}

	timeoutCtx, cancel := createTimeoutContext(ctx)
	defer cancel()

	user, err := s.r.GetUser(timeoutCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, formatUserResponse(user))
}

type (
	getAccountResponse struct {
		ID      int64   `json:"id"`
		UserID  int64   `json:"user_id"`
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}
)

// @Summary get account information
// @Description get account information
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path int true "account_id"
// @Success 200 {object} getAccountResponse
// @Router /accounts/{id} [get]
func (s *Server) getAccount(ctx *gin.Context) {
	id, err := getIDFromPath(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrIDInvalid, err.Error()))
		return
	}

	timeoutCtx, cancel := createTimeoutContext(ctx)
	defer cancel()

	acc, err := s.r.GetAccount(timeoutCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, formatAccountResponse(acc))
}

// @Summary list user account information
// @Description list account information
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {object} []getAccountResponse
// @Router /users/{id}/accounts [get]
func (s *Server) listUserAccounts(ctx *gin.Context) {
	id, err := getIDFromPath(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrIDInvalid, err.Error()))
		return
	}

	timeoutCtx, cancel := createTimeoutContext(ctx)
	defer cancel()

	accounts, err := s.r.ListAccounts(timeoutCtx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, formatListResponse(accounts, formatAccountResponse))
}
