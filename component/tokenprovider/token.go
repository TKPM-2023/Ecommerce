package tokenprovider

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewUnauthorized(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)

	ErrTokenExpired = common.NewUnauthorized(errors.New("token expired"),
		"token expired",
		"ErrTokenExpired",
	)
)
