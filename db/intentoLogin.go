package db

import (
	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin intenta hacer login buscando el email en la DB
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	//Verificar si el email del usuario esta en la DB
	usu, encontrado, _ := CheckExistingUser(email)

	if !encontrado {
		return usu, false
	}

	//Obtener ambas contraseñas
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	//comparar contraseñas
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
