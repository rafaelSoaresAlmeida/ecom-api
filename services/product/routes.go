package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rafaelSoaresAlmeida/ecom-api/types"
	"github.com/rafaelSoaresAlmeida/ecom-api/utils"
)

type Handler struct {
	store types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/product", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/product/{productId}", h.handleGetProductById).Methods(http.MethodGet)

	// admin routes
	//router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProduct()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productId"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.store.GetProductById(productId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if product.ID == 0 {
		utils.WriteJSON(w, http.StatusNotFound, nil)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct types.RegisterProductPayload

	if err := utils.ParseJson(r, &newProduct); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	if err := utils.Validate.Struct(newProduct); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	productDb, err := h.store.GetProductByName(newProduct.Name)

	if productDb.ID!= 0 && err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product with name %s already exists", newProduct.Name))
		return
	}

	err = h.store.CreateProduct(newProduct)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	utils.WriteJSON(w, http.StatusCreated, nil)
}