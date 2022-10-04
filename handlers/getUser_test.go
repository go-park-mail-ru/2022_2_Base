package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserErr400(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("GET", conf.PathGetUser, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		handler := http.HandlerFunc(userHandler.GetUser)

		log.Println(req.URL)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Code)
	})
}
