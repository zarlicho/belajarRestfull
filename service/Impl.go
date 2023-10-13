package service

import (
	"api/helper"
	"api/model/domain"
	"api/model/web"
	"api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ServiceImpl struct {
	Repository repository.Repository
	Db         *sql.DB
	Validate   *validator.Validate
}

func NewService(repo repository.Repository, Db *sql.DB, validate *validator.Validate) Service {
	return &ServiceImpl{Repository: repo, Db: Db, Validate: validate}
}

func (Service *ServiceImpl) Create(ctx context.Context, request web.CreateRequest) web.Response {
	err := Service.Validate.Struct(request)
	helper.PanicErrorIf(err)
	tx, err := Service.Db.Begin() //open database
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)

	category := domain.Datasiswa{
		Name:  request.Name,
		Kelas: request.Kelas,
	}

	category = Service.Repository.Save(ctx, tx, category)

	return helper.CategoryToResponse(category)
}

func (Service *ServiceImpl) Update(ctx context.Context, request web.UpdateRequest) web.Response {
	err := Service.Validate.Struct(request)
	helper.PanicErrorIf(err)
	tx, err := Service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)

	category, err := Service.Repository.FindByid(ctx, tx, request.Id)
	helper.PanicErrorIf(err)
	category.Name = request.Name
	category = Service.Repository.Update(ctx, tx, category)

	return helper.CategoryToResponse(category)
}

func (Service *ServiceImpl) Delete(ctx context.Context, datalistId int) {
	tx, err := Service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)
	category, err := Service.Repository.FindByid(ctx, tx, datalistId)
	helper.PanicErrorIf(err)
	Service.Repository.Delete(ctx, tx, category)
}

func (Service *ServiceImpl) FindByid(ctx context.Context, datalistId int) web.Response {
	tx, err := Service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)

	category, err := Service.Repository.FindByid(ctx, tx, datalistId)
	helper.PanicErrorIf(err)
	return helper.CategoryToResponse(category)
}

func (Service *ServiceImpl) FindAll(ctx context.Context) []web.Response {
	tx, err := Service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.TxErrHandle(tx)
	category := Service.Repository.FindAll(ctx, tx)
	helper.PanicErrorIf(err)
	var responseCategory []web.Response
	for _, categories := range category {
		responseCategory = append(responseCategory, helper.CategoryToResponse(categories))
	}
	return responseCategory
}
