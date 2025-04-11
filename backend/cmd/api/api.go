package main

import (
	"backend/internal/store"
	"net/http"
)

type application struct {
	Config config
	Store  *store.Storage
}

type config struct {
	addr string
}

func (app *application) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/v1/", app.corsMiddleware(http.StripPrefix("/v1", app.v1Routes())))

	return mux
}

func (app *application) v1Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.GetAllEmployees(w, r)
		case http.MethodPost:
			app.CreateEmployee(w, r)
		case http.MethodPut:
			app.UpdateEmployee(w, r)
		case http.MethodDelete:
			app.DeleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func (app *application) run(handler http.Handler) error {
	srv := &http.Server{
		Addr:    app.Config.addr,
		Handler: handler,
	}

	return srv.ListenAndServe()
}
