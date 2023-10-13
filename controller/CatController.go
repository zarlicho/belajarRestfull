package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Controller interface {
	Create(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindByid(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
}
