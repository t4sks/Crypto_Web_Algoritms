package httpserver

import (
	"Polibuis_Scytale/internal/httpserver/middleware"
	"net/http"
)

func New() *http.Server {
	mux := http.NewServeMux()
	registerRoutes(mux)

	handler := middlewareChain(
		mux,
		middleware.SecurityHeaders,
		middleware.RequestID,
		middleware.Logger,
		middleware.Recovery,
	)
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlePage)
	mux.HandleFunc("/api", handleApi)
	mux.Handle(
		"/static/",
		cacheControlStatic(
			http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))))
}

func middlewareChain(next http.Handler, middlewares ...func(handler http.Handler) http.Handler) http.Handler {
	wrapped := next
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}
	return wrapped
}

func cacheControlStatic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=300")
		next.ServeHTTP(w, r)
	})

}
