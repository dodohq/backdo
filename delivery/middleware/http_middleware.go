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
		} else if claims["is_admin"] != true {
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
