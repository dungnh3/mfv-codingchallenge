// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package store

import (
	"context"
)

type Querier interface {
	Ping(ctx context.Context) error
}

var _ Querier = (*Queries)(nil)
