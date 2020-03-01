package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/otus/calendar/configs"
	"go.uber.org/zap"
	"net/http"
)

// NewHTTPServer get http server instance.
func NewHTTPServer(cfg configs.ServerConfig) *http.Server {
	defer zap.L().Info(fmt.Sprintf("Http server %s is started!", cfg.Addr))
	router := mux.NewRouter()
	router.HandleFunc("/hello", LoggerMiddleware(Hello)).Methods("GET")

	return &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
}

// LoggerMiddleware logger middleware to http endpoints.
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		zap.L().Info(fmt.Sprintf("Request %s | %s%s | Client:%s",
			request.Method, request.Host, request.URL, request.RemoteAddr))
		next(writer, request)
	}
}

// Hello response hello string.
func Hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello")); err != nil {
		zap.L().Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
