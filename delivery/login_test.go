package delivery

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"strings"
	"testing"

	baseErrors "serv/errors"

	"github.com/stretchr/testify/assert"
)

var casesLogin = []struct {
	data     map[string]string
	wantCode int
	err      error
}{
	{map[string]string{"email": "s", "username": "art", "password": "string"}, 401, baseErrors.ErrUnauthorized401},
	{map[string]string{"email": "string", "username": "art", "password": "s"}, 401, baseErrors.ErrUnauthorized401},
}

func TestLogin(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		data, _ := json.Marshal(map[string]string{"email": "art@art", "username": "art", "password": "art"})
		req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		userHandler.Login(rr, req)
		assert.Equal(t, 201, rr.Code)
	})

}

func TestLoginErrors(t *testing.T) {
	for _, c := range casesLogin {
		t.Run("tests", func(t *testing.T) {
			data, _ := json.Marshal(c.data)
			req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			userHandler := NewUserHandler()
			userHandler.Login(rr, req)
			assert.Equal(t, c.wantCode, rr.Code)
		})
	}
}

func TestLoginErr400(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("POST", conf.PathSignUp, strings.NewReader("ASDSAD"))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		userHandler.Login(rr, req)
		assert.Equal(t, 400, rr.Code)
	})
}
