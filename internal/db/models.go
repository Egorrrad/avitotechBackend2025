package db

import (
	"avitotechBackend2025/internal/dto"
	"github.com/oapi-codegen/runtime/types"
)

type UserModel struct {
	dto.User
	Password string
}

type PvzModel struct {
	dto.PVZ
	createdBy types.UUID
}

type ReceptionModel struct {
	dto.Reception
	createdBy types.UUID
}

type ProductModel struct {
	dto.Product
}
