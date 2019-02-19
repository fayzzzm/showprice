package server

import (
	"billing/site/configs"
	"billing/site/router"
	"net/http"
	"time"
)

func Init() (err error) {
	appServer := &http.Server{
		Addr:         configs.Get().Server.Port,
		Handler:      router.Get(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	err = appServer.ListenAndServe()

	return
}
