package helper

import (
	"api/model/domain"
	"api/model/web"
)

func CategoryToResponse(datalist domain.Datasiswa) web.Response {
	return web.Response{
		Id:    datalist.Id,
		Name:  datalist.Name,
		Kelas: datalist.Kelas,
	}
}
