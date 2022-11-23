package handlers

import (
	"burmachine/TestSolution/internal/auth"
	"fmt"
	"net/http"
)

type AuthMiddlwareData struct {
	Data auth.AuthData
}

func (a *AuthMiddlwareData) BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth() // Basic auth
		fmt.Printf("User: %s; Pass: %s; OK: %v\n", username, password, ok)
		if !ok || !a.Data.CheckUserNameAndPassword(username, password) {
			writer.Header().Set(`WWW-Authenticate`, `Basic auth="Account invalid"`)
			writer.WriteHeader(401)
			writer.Write([]byte("Unauthorised!\n"))
		} else {
			handler(writer, request)
		}
	}
}
