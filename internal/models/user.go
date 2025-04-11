package models

import "github.com/oapi-codegen/runtime/types"

type User struct {
	Email    types.Email
	Password string
	Role     string
}
