package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindByid(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)

	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
