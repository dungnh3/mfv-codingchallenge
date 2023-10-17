package services

import (
	"net/http"

	errorz "github.com/dungnh3/mfv-codingchallenge/internal/errors"
	"github.com/dungnh3/mfv-codingchallenge/internal/models"
	"github.com/dungnh3/mfv-codingchallenge/internal/models/store"
	"github.com/gin-gonic/gin"
)

type (
	RegisterUserRequest struct {
		Name string `json:"name" binding:"required"`
	}

	RegisterUserResponse struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

// @Summary register new user
// @Description register new user
// @Tags Users
// @Accept json
// @Produce json
// @Param register body RegisterUserRequest true "register a new user"
// @Success 200 {object} RegisterUserResponse
// @Router /users/register [post]
func (s *Server) registerUser(ctx *gin.Context) {
	var request RegisterUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrBadParamInput, err.Error()))
		return
	}

	timeoutCtx, cancel := createTimeoutContext(ctx)
	defer cancel()

	user := &models.User{
		User: &store.User{
			Name:   request.Name,
			Status: models.UserStatusActive.String(),
		},
	}

	if err := s.r.UpsertUser(timeoutCtx, user); err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(err))
		return
	}
	response := RegisterUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
	ctx.JSON(http.StatusOK, response)
}

type (
	GetUserResponse struct {
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
// @Success 200 {object} GetUserResponse
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
	CreateAccountRequest struct {
		Name string `json:"name" binding:"required"`
	}

	CreateAccountResponse struct {
		ID     int64  `json:"id"`
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
	}
)

// @Summary create new account of user
// @Description create new account of user
// @Tags Users
// @Accept json
// @Produce json
// @Param register body CreateAccountRequest true "create a new account"
// @Success 200 {object} CreateAccountResponse
// @Router /users/:id/accounts [post]
func (s *Server) createAccount(ctx *gin.Context) {
	id, err := getIDFromPath(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrIDInvalid, err.Error()))
		return
	}

	var request CreateAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(errorz.ErrBadParamInput, err.Error()))
		return
	}

	timeoutCtx, cancel := createTimeoutContext(ctx)
	defer cancel()

	acc := &models.UserAccount{
		UserAccount: &store.UserAccount{
			UserID: int64(id),
			Name:   request.Name,
			Status: models.UserStatusActive.String(),
		},
	}

	if err := s.r.UpsertUserAccount(timeoutCtx, acc); err != nil {
		ctx.JSON(http.StatusBadRequest, errorz.NewErrResponse(err))
		return
	}
	response := CreateAccountResponse{
		ID:     acc.ID,
		UserID: acc.UserID,
		Name:   acc.Name,
	}
	ctx.JSON(http.StatusOK, response)
}

type (
	GetAccountResponse struct {
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
// @Success 200 {object} GetAccountResponse
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
// @Success 200 {object} []GetAccountResponse
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
