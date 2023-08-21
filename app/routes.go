package app

import (
	"encoding/json"
	"fmt"
	"github.com/TotemFi/totem-bridge-offchain/util"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

func (app *App) Serve() {
	router := mux.NewRouter()
	// We use our custom CORS Middleware
	router.Use(CORS)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	var api = router.PathPrefix("/api/v1").Subrouter()

	var adminApi = api.PathPrefix("/app").Subrouter()
	adminApi.HandleFunc("/", app.Endpoints).Methods("GET")
	adminApi.HandleFunc("/healthz", app.Healthz).Methods("GET")
	adminApi.HandleFunc("/update_swap_pair", app.UpdateSwapPairHandler).Methods("PUT")
	adminApi.HandleFunc("/withdraw_token", app.WithdrawToken).Methods("POST")
	adminApi.HandleFunc("/retry_failed_swaps", app.RetryFailedSwaps).Methods("POST")

	// user routes
	api.HandleFunc("/bridge/price", app.makePriceHandler).Methods("GET")
	api.HandleFunc("/bridge/price/signature", app.makePriceSignatureHandler).Methods("GET")
	api.HandleFunc("/bridge/pegin/state", app.peginStateHandler).Methods("GET")

	listenAddr := DefaultListenAddr
	if app.cfg.ServerConfig.ListenAddr != "" {
		listenAddr = app.cfg.ServerConfig.ListenAddr
	}
	//handlers := cors.Default().Handler(router)
	//router.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))

	srv := &http.Server{
		Handler:      router,
		Addr:         listenAddr,
		WriteTimeout: 90 * time.Second,
		ReadTimeout:  90 * time.Second,
	}

	util.Logger.Infof("start app server at %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("start app server error, err=%s", err.Error()))
	}
}

func (app *App) Endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: []string{
			"/update_swap_pair",
			"/healthz",
		},
	}
	jsonBytes, err := json.MarshalIndent(endpoints, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		util.Logger.Errorf("write response error, err=%s", err.Error())
	}
}

func (app *App) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//origin := r.Header.Get("Origin")
		// add header Access-Control-Allow-Origin
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // TODO
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,  Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
