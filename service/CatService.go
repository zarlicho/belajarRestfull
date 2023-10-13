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
}
