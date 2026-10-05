package utils

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	SendJson(map[string]string{"status": "ok"}, w, r)
}
