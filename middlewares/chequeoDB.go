package middlewares

import (
	"net/http"

	"github.com/Marlos-Rodriguez/Twitter-Clon-Back/db"
)

//ChequeoDB Middleware que comprueba la conexion de la DB
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
