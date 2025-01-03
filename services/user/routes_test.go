package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rafaelSoaresAlmeida/ecom-api/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload {
			FirstName: "user",
			LastName: "123",
			Email: "",
			Password: "asd",
		}

		rr := makeUserRegistryEndpointCall(payload,  handler)
		validateResponse(rr , http.StatusBadRequest, t)
	})

		t.Run("should create user with success", func(t *testing.T) {
		payload := types.RegisterUserPayload {
			FirstName: "tchucotchuco",
			LastName: "supersuper",
			Email: "supertchuco@toco.com",
			Password: "abc235123",
		}

		rr := makeUserRegistryEndpointCall(payload,  handler)

		validateResponse(rr , http.StatusCreated, t)
	})
}

type mockUserStore struct{}

/* func (m *mockUserStore) UpdateUser(u types.User) error {
	return nil
} */

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return &types.User{}, nil
}

func makeUserRegistryEndpointCall(payload types.RegisterUserPayload,  handler *Handler) *httptest.ResponseRecorder {
	marshalled, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/register", handler.handleRegister)
	router.ServeHTTP(rr, req)
	return rr
}

func validateResponse(rr *httptest.ResponseRecorder, httpStatusResult int, t *testing.T) {
			if rr.Code != httpStatusResult {
			t.Errorf("expected status code %d, got %d", httpStatusResult, rr.Code)
		}
} 