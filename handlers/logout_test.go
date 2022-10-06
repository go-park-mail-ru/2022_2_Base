package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogout(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		data, _ := json.Marshal(map[string]string{"username": "art", "password": "art"})
		req, err := http.NewRequest("POST", conf.PathLogin, strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		productHandler := NewUserHandler()
		handler := http.HandlerFunc(productHandler.Login)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		req, err = http.NewRequest("DELETE", conf.PathLogOut, nil)
		req.AddCookie(rr.Result().Cookies()[0])
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()

		handler = http.HandlerFunc(productHandler.Logout)

		handler.ServeHTTP(rr, req)
		log.Println(rr.Body.String())
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestLogoutErr401(t *testing.T) {

	t.Run("tests", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", conf.PathLogOut, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		handler := http.HandlerFunc(userHandler.Logout)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, 401, rr.Code)
	})

}
