package handler

import (
	"net/http"

	"github.com/andsholinka/Digitalent-Kominfo_Go-Microservice/menu-service/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
