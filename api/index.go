package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
)

var (
	RestApi config.RestConfig
	initErr error
	once    sync.Once
)

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		log.Println("Initializing RestApi...")
		RestApi, initErr = config.NewRest()
		if initErr != nil {
			log.Printf("RestApi initialization failed: %v", initErr)
		} else {
			log.Println("RestApi initialized successfully")
		}
	})

	if initErr != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Infrastructure Error: %v\n", initErr)
		return
	}

	RestApi.GetServer().ServeHTTP(w, r)
}
