package api

import (
	"net/http"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
)

var (
	RestApi config.RestConfig
)

func init() {
	RestApi = config.NewRest()
}

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	RestApi.GetServer().ServeHTTP(w, r)
}
