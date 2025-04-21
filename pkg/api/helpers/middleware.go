package helpers

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Middleware = func(httprouter.Handle) httprouter.Handle

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		next(w, r, ps)
		duration := time.Since(start).String()
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, duration)
	}
}

func ApplyMiddleware(handler httprouter.Handle, middlewares []Middleware) httprouter.Handle {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
