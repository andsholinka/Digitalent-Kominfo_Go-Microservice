package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/config"
	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/utils"
)

type AuthHandler struct {
	Config config.Auth
}

func (handler *AuthHandler) ValidateAdmin(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest("POST", handler.Config.Host+"/admin-auth", nil)
		if err != nil {
			utils.WrapAPIError(w, r, "failed to create request : "+err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header
		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, "validate auth failed : "+err.Error(), http.StatusInternalServerError)
			return
		}
		// defer authResponse.Body.Close()

		responBody, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var responData map[string]interface{}
		err = json.Unmarshal(responBody, &responData)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, "Invalid auth", authResponse.StatusCode)
			return
		}

		nextHandler(w, r)
	}
}
