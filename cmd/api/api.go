package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelSoaresAlmeida/ecom-api/services/cart"
	"github.com/rafaelSoaresAlmeida/ecom-api/services/order"
	"github.com/rafaelSoaresAlmeida/ecom-api/services/product"
	"github.com/rafaelSoaresAlmeida/ecom-api/services/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer (addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db: db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subRouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subRouter)

	// Serve static files ????
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Up and Running on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}