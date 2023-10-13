package services

import (
	"context"
	"strconv"
	"time"

	"github.com/dungnh3/mfv-codingchallenge/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	idParam = "id"

	Timeout5s = 5 * time.Second
)

func getIDFromPath(ctx *gin.Context) (int, error) {
	return strconv.Atoi(ctx.Param(idParam))
}

func createTimeoutContext(ctx *gin.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx.Request.Context(), Timeout5s)
}

func formatListResponse[V, T any](arr []V, fn func(V) T) (results []T) {
	for _, v := range arr {
		results = append(results, fn(v))
	}
	return
}

func formatAccountResponse(acc *models.UserAccount) getAccountResponse {
	return getAccountResponse{
		ID:      acc.ID,
		UserID:  acc.UserID,
		Name:    acc.Name,
		Balance: acc.Balance,
	}
}

func formatUserResponse(u *models.User) getUserResponse {
	return getUserResponse{
		ID:         u.ID,
		Name:       u.Name,
		AccountIDs: u.AccountIDs,
	}
}
