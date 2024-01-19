package utils

import (
	"context"
	"user-app/internal/consts"
	"user-app/internal/dto/input"

	"github.com/google/uuid"
)

func GetCreateUserRequestFromCtx(ctx context.Context) input.CreateUserRequest {
	if val, ok := ctx.Value(consts.ContextKeyCreateUserRequest).(input.CreateUserRequest); ok {
		return val
	}
	return input.CreateUserRequest{}
}

func GetUpdateUserRequestFromCtx(ctx context.Context) input.UpdateUserRequest {
	if val, ok := ctx.Value(consts.ContextKeyUpdateUserRequest).(input.UpdateUserRequest); ok {
		return val
	}
	return input.UpdateUserRequest{}
}

func GetUserIDFromCtx(ctx context.Context) uuid.UUID {
	if val, ok := ctx.Value(consts.UserId).(uuid.UUID); ok {
		return val
	}
	return uuid.Nil
}