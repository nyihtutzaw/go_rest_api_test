package authservice

import (
	"encoding/json"
	"net/http"
	"rest_api_test/models"
)

//Middleware func
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token, ok := ExtractToken(r)
		if ok {
			valid := CheckToken(token)
			if valid {
				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(models.ResponseType{
					Message: "Token is invalid",
				})
			}

		} else {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(models.ResponseType{
				Message: "Token is required",
			})
		}

	})
}
