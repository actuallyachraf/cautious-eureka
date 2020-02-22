package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

// Service wraps the API service backend.
type Service struct {
	Router *mux.Router
}

// commonModdleWare is a re-usable middleware to specify content types
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// NewService creates a new instance of the API service
func NewService() *Service {

	return &Service{
		Router: mux.NewRouter(),
	}

}

// SetupService registers the handlers.
func (serv *Service) SetupService() {
	serv.Router.Use(commonMiddleware)
	serv.Router.PathPrefix("/api").Methods("GET").Path(TrendingRoute).HandlerFunc(TrendingHandler)
	serv.Router.PathPrefix("/api").Methods("GET").Path(LanguageRoute).HandlerFunc(LanguageHandler)
}

// Start the service and logging it launches two go routines
// one for serving the API and one of listening to system interrupts.
func (serv *Service) Start() {

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		log.Println("trendy is listening on port:", Port)
		errChan <- http.ListenAndServe(Port, serv.Router)
	}()

	log.Fatalln(<-errChan)
}
