package api

import (
	"net/http"
	"sync"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
)

var (
	RestApi config.RestConfig
	once    sync.Once
)

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	once.Do(func() {
		RestApi, err = config.NewRest()
	})

	if err != nil {
		http.Error(w, "Infrastructure Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	RestApi.GetServer().ServeHTTP(w, r)
}
