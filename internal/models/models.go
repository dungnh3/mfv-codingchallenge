package models

import "github.com/dungnh3/mfv-codingchallenge/internal/models/store"

type (
	AccountIDs []int64
	User       struct {
		*store.User
		AccountIDs AccountIDs `json:"account_ids"`
	}

	UserAccount struct {
		*store.UserAccount
	}
)
