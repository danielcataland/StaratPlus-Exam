package utils

import (
	"log"
	"stratplusapi/internal/models"
)

func ValidateRequest(u models.User, login bool) string {

	log.Println("[INFO]: Validating body request")

	if login {
		if validateEmail(u) != "El campo email está vacío" ||
			validateUserName(u) != "El campo username está vacío" {
			if res := validatePwd(u); res != "" {
				return res
			}

			if res := validatePhone(u); res != "" {
				return res
			}
			return ""
		}
	} else {
		if res := validateEmail(u); res != "" {
			return res
		}
		if res := validateUserName(u); res != "" {
			return res
		}
	}

	if res := validatePwd(u); res != "" {
		return res
	}

	if res := validatePhone(u); res != "" {
		return res
	}

	log.Println("[INFO]: Body request validated successfully")
	return ""
}

func validateEmail(u models.User) string {
	if u.Email == "" {
		return "El campo email está vacío"
	}

	if email := MatchWord(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, u.Email); !email {
		return "El campo e-mail no tiene un formato válido"
	}

	return ""
}

func validatePwd(u models.User) string {
	if u.Password == "" {
		return "El campo password está vacío"
	}

	if pwd := MatchWord(`^.{6,12}$`, u.Password); !pwd {
		return "El campo password debe ser de una longitud mayor a 6 y menor a 12"
	}

	if MatchWord(`^(?=.*\d)(?=.*[a-zA-Z])(?=.*[A-Z])(?=.*[-\@\$\&])(?=.*[a-zA-Z]).{6,12}$`,
		u.Password) {
		return "El campo de password no tiene un formato válido"
	}

	return ""
}

func validateUserName(u models.User) string {
	if u.UserName == "" {
		return "El campo username está vacío"
	}

	if !MatchWord(`^[a-zA-Z0-9]{3,24}$`, u.UserName) {
		return "El campo username deben ser números, letras y menor a 24 caracteres"
	}

	return ""
}

func validatePhone(u models.User) string {
	if u.Phone == "" {
		return "El campo phone está vacío"
	}

	if phone := MatchWord(`^[0-9]{10}$`, u.Phone); !phone {
		return "La longitud del teléfono debe ser igual a 10 dígitos"
	}
	return ""
}
