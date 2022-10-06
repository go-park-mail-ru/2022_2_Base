package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesLogin = []struct {
	data     map[string]string
	wantCode int
	err      error
}{
	{map[string]string{"username": "art", "password": "art"}, 201, nil},
	{map[string]string{"username": "s", "password": "string"}, 400, nil},
	{map[string]string{"username": "string", "password": "s"}, 400, nil},
	{map[string]string{"sdads": "d"}, 400, nil},
}

func TestLogin(t *testing.T) {
	for _, c := range casesLogin {
		t.Run("tests", func(t *testing.T) {
			data, _ := json.Marshal(c.data)
			req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			userHandler := NewUserHandler()
			handler := http.HandlerFunc(userHandler.Login)

			handler.ServeHTTP(rr, req)
			assert.Equal(t, c.wantCode, rr.Code)
		})
	}
}

func TestLoginErr400(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader("ASDSAD"))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		handler := http.HandlerFunc(userHandler.Login)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Code)
	})
}
