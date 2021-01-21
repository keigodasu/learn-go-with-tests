package http_handlers_revisited

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type MockUserService struct {
	RegisterFunc    func(user User) (string, error)
	UsersRegistered []User
}

func (m *MockUserService) Register(user User) (insertedID string, err error) {
	m.UsersRegistered = append(m.UsersRegistered, user)
	return m.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
	t.Run("can register valid users", func(t *testing.T) {
		user := User{Name: "CJ"}
		expectedInsertedId := "whatever"

		service := MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return expectedInsertedId, nil
			},
		}
		server := NewUserServer(&service)

		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)
		assertStatus(t, res.Code, http.StatusCreated)

		if res.Body.String() != expectedInsertedId {
			t.Errorf("expected body of %q but got %q", res.Body.String(), expectedInsertedId)
		}

		if len(service.UsersRegistered) != 1 {
			t.Fatalf("expected 1 user added but got %d", len(service.UsersRegistered))
		}

		if !reflect.DeepEqual(service.UsersRegistered[0], user) {
			t.Errorf("the user registered %+v was not what was expected %+v", service.UsersRegistered[0], user)
		}

	})

	t.Run("returns 400 bad request if body is not valid user JSON", func(t *testing.T) {
		server := NewUserServer(nil)

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("valid string"))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)
		assertStatus(t, res.Code, http.StatusBadRequest)
	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {
		user := User{Name: "CJ"}

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
				return "", errors.New("couldnt add new suer")
			},
		}
		server := NewUserServer(service)
		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res.Code, http.StatusInternalServerError)
	})
}

func assertStatus(t *testing.T, code int, expected int) {
	if code != expected {
		t.Error()
	}
}

func userToJSON(user User) io.Reader {
	json, _ := json.Marshal(user)
	return strings.NewReader(string(json))
}
