package handler

import (
	"encoding/json"
	"log"
	"net/http"
	models "stratplusapi/internal/models"
	"stratplusapi/internal/security"
	ut "stratplusapi/internal/utils"
)

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("[INFO]: Login handler started")

	var usrModel models.User
	err := json.NewDecoder(req.Body).Decode(&usrModel)
	if err != nil {
		log.Println("[ERROR]: Error unmarshaling body request: ", err)
		res := models.CreateResponse(rw, http.StatusBadRequest, "El body de la petición es erróneo", "")
		res.SendResult()
		return
	}

	valRes := ut.ValidateRequest(usrModel, true)
	if valRes != "" {
		res := models.CreateResponse(rw, http.StatusBadRequest, valRes, "")
		res.SendResult()
		return
	}

	log.Println("[INFO]: Login user")
	usrExi, inf := loginUser(usrModel)
	if inf == "Internal server error" {
		res := models.CreateResponse(rw, http.StatusInternalServerError, "Internal server error", "")
		res.SendResult()
		return
	} else if inf == "NE" {
		res := models.CreateResponse(rw, http.StatusUnauthorized, "Usuario o contraseña incorrectos", "")
		res.SendResult()
		return
	}

	token, err := security.CreateToken(usrExi.UserName, usrExi.Phone, usrExi.Email)
	if err != nil {
		log.Println("[ERROR]: Error generating token: ", err)
		res := models.CreateResponse(rw, http.StatusInternalServerError, "Internal server error", "")
		res.SendResult()
		return
	}

	log.Println("[INFO]: Token created successfully")
	res := models.CreateResponse(rw, http.StatusOK, "", token)
	res.SendResult()

}

func loginUser(us models.User) (models.User, string) {

	usr := new(models.User)

	res, err := usr.GetUser(us, true)
	if err == "Internal server error" {
		log.Println("[ERROR]: Error to get user by username or email and password: ", err)
		return models.User{}, err
	} else if err == "exist" && res.Email != "" {
		log.Println("[INFO]: User exist DB")
		return res, ""
	} else if err == "exist" && res.Password == "" {
		log.Println("[INFO]: User not exist")
		return res, "NE"
	}

	return res, ""
}
