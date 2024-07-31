package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newRouter() *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/youtrack/channel/stats", getChannelStats())

	return mux
}

func getChannelStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("response!"))
	}
}

func main() {
	srv := &http.Server{
		Addr:    ":5000",
		Handler: newRouter(),
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fatal http server failed to start: %v", err)
		}
	}
}
