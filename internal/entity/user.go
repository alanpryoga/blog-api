package entity

import "errors"

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type User struct {
	ID             int32
	Email          string
	HashedPassword string
}
