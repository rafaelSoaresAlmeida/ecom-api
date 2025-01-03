package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rafaelSoaresAlmeida/ecom-api/config"
	"github.com/rafaelSoaresAlmeida/ecom-api/services/auth"
	"github.com/rafaelSoaresAlmeida/ecom-api/types"
	"github.com/rafaelSoaresAlmeida/ecom-api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store : store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var userPayload types.LoginUserPayload

	if err := utils.ParseJson(r, &userPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	if err := utils.Validate.Struct(userPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	dbUser, err := h.store.GetUserByEmail(userPayload.Email)

	if err != nil  {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid credentials"))
		return
	}

	if !auth.ComparePasswords(dbUser.Password, []byte(userPayload.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid credentials"))
		return
	}

	secret := []byte(config.Envs.JwtSecret)
	token, err := auth.CreateJwt(secret, dbUser.ID) 

	if err != nil  {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string] string {"token": token})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload

	if err := utils.ParseJson(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	userDb, err := h.store.GetUserByEmail(user.Email)

	if userDb.ID!= 0 && err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	utils.WriteJSON(w, http.StatusCreated, nil)

}