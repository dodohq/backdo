package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dodohq/backdo/delivery/helper"
	"github.com/dodohq/backdo/lib/jwt"
	"github.com/dodohq/backdo/models"
	"github.com/julienschmidt/httprouter"
)

// AdminAuthy middleware to check if this user is logged in as admin
func AdminAuthy(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			err := models.NewErrorUnauthorized("No Token Present")
			helper.RenderErr(w, err)
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			err := models.NewErrorInternalServer(err.Error())
			helper.RenderErr(w, err)
			return
		} else if claims["role"] != jwt.AdminType {
			err := models.NewErrorUnauthorized("No Admin Access")
			helper.RenderErr(w, err)
			return
		} else if claims["exp"].(float64) <= float64(time.Now().Unix()) {
			err := models.NewErrorUnauthorized("Token Expired")
			helper.RenderErr(w, err)
			return
		}

		jsonBytes, err := json.Marshal(claims)
		r.Header.Set("User", string(jsonBytes))
		handler(w, r, ps)
	}
}

// UserAuthy middleware to check if this user is logged in
func UserAuthy(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			err := models.NewErrorUnauthorized("No Token Present")
			helper.RenderErr(w, err)
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			err := models.NewErrorInternalServer(err.Error())
			helper.RenderErr(w, err)
			return
		} else if claims["role"] != jwt.UserType {
			helper.RenderErr(w, models.NewErrorUnauthorized("No Staff Access"))
			return
		} else if claims["exp"].(float64) <= float64(time.Now().Unix()) {
			err := models.NewErrorUnauthorized("Token Expired")
			helper.RenderErr(w, err)
			return
		}

		jsonBytes, err := json.Marshal(claims)
		r.Header.Set("User", string(jsonBytes))
		handler(w, r, ps)
	}
}

// DriverAuthy middleware to check if this driver is logged in
func DriverAuthy(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			helper.RenderErr(w, models.NewErrorUnauthorized("No Token Present"))
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			helper.RenderErr(w, models.NewErrorInternalServer(err.Error()))
			return
		} else if claims["role"] != jwt.DriverType {
			helper.RenderErr(w, models.NewErrorUnauthorized("No Driver Access"))
			return
		} else if claims["exp"].(float64) <= float64(time.Now().Unix()) {
			helper.RenderErr(w, models.NewErrorUnauthorized("Token Expired"))
			return
		}

		jsonBytes, err := json.Marshal(claims)
		r.Header.Set("User", string(jsonBytes))
		handler(w, r, ps)
	}
}
