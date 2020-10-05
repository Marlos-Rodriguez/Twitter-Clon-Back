package models

//ResponseLogin tiene el token que se devuelve con el login
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
