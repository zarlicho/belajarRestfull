package service

import (
	"api/model/web"
	"context"
)

type Service interface {
	Create(ctx context.Context, requests web.CreateRequest) web.Response
	Update(ctx context.Context, requests web.UpdateRequest) web.Response
	Delete(ctx context.Context, datalistId int)
	FindByid(ctx context.Context, datalistId int) web.Response
	FindAll(ctx context.Context) []web.Response

	Register(ctx context.Context, requests web.RegisRequest)
	Login(ctx context.Context, requests web.LoginRequest) (string, error)
}
