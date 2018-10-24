package route

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopx.io/gopx/internal/app/gopx/handler"
	"gopx.io/gopx/pkg/log"
)

// Router registers the routes.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)

	r.Use(loggingMiddleware)

	r.Path("/").
		Methods("GET").
		HandlerFunc(handler.IndexGET)

	return r
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("%s %s", strings.ToUpper(r.Method), r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
