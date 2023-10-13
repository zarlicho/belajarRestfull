package controller

import (
	"api/helper"
	"api/model/web"
	"api/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ControllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &ControllerImpl{Service: Service}
}

func (Controller *ControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//receiving body data atau json
	decoder := json.NewDecoder(request.Body)
	CreateRequest := web.CreateRequest{}
	err := decoder.Decode(&CreateRequest)
	helper.PanicErrorIf(err)

	categoryResponse := Controller.Service.Create(request.Context(), CreateRequest)
	//write response
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)

}

func (Controller *ControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	UpdateRquest := web.UpdateRequest{}
	err := decoder.Decode(&UpdateRquest)
	helper.PanicErrorIf(err)

	paramsid := params.ByName("id")
	id, err := strconv.Atoi(paramsid)
	helper.PanicErrorIf(err)
	UpdateRquest.Id = id

	categoryResponse := Controller.Service.Update(request.Context(), UpdateRquest)
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)

}

func (Controller *ControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramsid := params.ByName("id")
	id, err := strconv.Atoi(paramsid)
	helper.PanicErrorIf(err)

	Controller.Service.Delete(request.Context(), id)
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)
}
func (Controller *ControllerImpl) FindByid(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramsid := params.ByName("id")
	id, err := strconv.Atoi(paramsid)
	helper.PanicErrorIf(err)
	categoryResponse := Controller.Service.FindByid(request.Context(), id)

	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)
}
func (Controller *ControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := Controller.Service.FindAll(request.Context())
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicErrorIf(err)
}
