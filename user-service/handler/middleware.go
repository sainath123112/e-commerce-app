package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sainath/e-commerce-app/common/pkg"
	"github.com/sainath/e-commerce-app/user-service/model"
	"github.com/sainath/e-commerce-app/user-service/service"
	"net/http"
	"strconv"
)

func Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		vars := mux.Vars(r)
		userid, err := strconv.Atoi(vars["userid"])
		w.Header().Set("Content-Type", "application/json")
		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No header authorization found", ErrorString: "Unauthorized"})
			return
		}
		token := authorization[len("Bearer "):]
		isValidToken, err := pkg.ValidateToken(token)
		if !isValidToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "Invalid Token", ErrorString: err.Error()})
			return
		}
		emailFromToken, err := pkg.GetUsername(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: emailFromToken, ErrorString: err.Error()})
			return
		}
		emailFromUserId, err := service.GetUserEmail(userid)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "No user found for id: " + vars["userid"], ErrorString: err.Error()})
			return
		}
		if emailFromUserId != emailFromToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.ErrorResponse{Message: "token is not belongs to this user with user id: " + vars["userid"], ErrorString: "Token mismatch"})
			return
		}
		handler.ServeHTTP(w, r)
	}
}
