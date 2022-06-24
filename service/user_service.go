package service

import (
	"context"

	"github.com/vincen320/user-service/model/appservice/authservice"
	"github.com/vincen320/user-service/model/web"
)

type UserService interface {
	Create(ctx context.Context, userCreate web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, userUpdate web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int) bool
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	UpdatePatch(ctx context.Context, userUpdate web.UserUpdatePatchRequest) web.UserResponse
	FindByUsername(ctx context.Context, username string) authservice.UserServiceLoginResponse
}
