package handler

import (
	"net/http"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/utils"
)

func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "respecker" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}
