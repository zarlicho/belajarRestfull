package controller

import (
	"api/helper"
	"api/model/web"
	"api/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
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

func (Controller *ControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//receiving body data atau json
	decoder := json.NewDecoder(request.Body)
	RegisRequest := web.RegisRequest{}
	err := decoder.Decode(&RegisRequest)
	helper.PanicErrorIf(err)

	Controller.Service.Register(request.Context(), RegisRequest)
	//write response
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)

}
func (Controller *ControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//receiving body data atau json
	decoder := json.NewDecoder(request.Body)
	LoginRequest := web.LoginRequest{}
	err := decoder.Decode(&LoginRequest)
	helper.PanicErrorIf(err)
	fmt.Println(LoginRequest)
	token, errs := Controller.Service.Login(request.Context(), LoginRequest)
	if errs != nil {
		errors.New("error at controller pattern")
	}
	// Membuat cookie token
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,                               // Token JWT yang diterima dari Service
		Expires:  time.Now().Add(30 * 24 * time.Hour), // Atur durasi cookie sesuai kebutuhan
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // Atur SameSite sesuai kebutuhan
		Secure:   true,                 // Hanya diaktifkan saat menggunakan HTTPS
	}

	// Menambahkan cookie ke respons
	http.SetCookie(writer, cookie)
	//write response
	webResponse := web.ResponseWeb{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicErrorIf(err)

}
