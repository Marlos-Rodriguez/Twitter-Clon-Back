package middlewares

import (
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/routers"
)

//ValidJWT permite validar el JWT que nos viene en la peticion
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error en el token! "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
