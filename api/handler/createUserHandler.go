package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	models "stratplusapi/internal/models"
	ut "stratplusapi/internal/utils"
)

func CreateUserHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("[INFO]: Create user handler started")

	var usrModel models.User
	err := json.NewDecoder(req.Body).Decode(&usrModel)
	if err != nil {
		log.Println("[ERROR]: Error unmarshaling body request: ", err)
		res := models.CreateResponse(rw, http.StatusBadRequest, "El body de la petición es erróneo", "")
		res.SendResult()
		return
	}

	valRes := ut.ValidateRequest(usrModel, false)
	if valRes != "" {
		res := models.CreateResponse(rw, http.StatusBadRequest, valRes, "")
		res.SendResult()
		return
	}

	log.Println("[INFO]: Getting user by phone and email")
	usrExi, inf := UserExist(usrModel)
	if usrExi && inf == "" {
		res := models.CreateResponse(rw, http.StatusNotFound, "El correo/teléfono ya se encuentra registrado", "")
		res.SendResult()
		return
	}
	if inf == "Internal server error" {
		res := models.CreateResponse(rw, http.StatusInternalServerError, "Internal server error", "")
		res.SendResult()
		return
	}
	result := models.User.CreateUsr(usrModel)

	if result != nil {
		res := models.CreateResponse(rw, http.StatusInternalServerError, "Internal server error", "")
		res.SendResult()
		return
	} else {
		res := models.CreateResponse(rw, http.StatusCreated, "Usuario creado", "")
		res.SendResult()
	}
}

func UserExist(us models.User) (bool, string) {

	usr := new(models.User)

	res, err := usr.GetUser(us, false)
	if err == "Internal server error" {
		fmt.Println("[ERROR]: Error to get user by phone or email: ", err)
		return true, err
	} else if err == "exist" && res.Email != "" {
		fmt.Println("[ERROR]: User exist in DB")
		return true, ""
	}

	fmt.Println(usr.Email)
	return false, ""
}
