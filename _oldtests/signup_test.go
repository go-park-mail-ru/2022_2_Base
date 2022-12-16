package delivery

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	conf "serv/config"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesSignUp = []struct {
	data     map[string]string
	wantCode int
	err      error
}{
	{map[string]string{"email": "art@art22", "username": "string", "password": "111"}, 401, nil},
	{map[string]string{"email": "art@art", "username": "art", "password": "111111"}, 409, nil},
	{map[string]string{"username": "artart", "password": "111111"}, 401, nil},
	{map[string]string{"email": "123", "username": "artart", "password": "111111"}, 401, nil},
}

func TestSignUpErrors(t *testing.T) {
	for _, c := range casesSignUp {
		t.Run("tests", func(t *testing.T) {
			data, _ := json.Marshal(c.data)
			req, err := http.NewRequest("POST", conf.PathSignUp, strings.NewReader(string(data)))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			userHandler := NewUserHandler()
			userHandler.SignUp(rr, req)
			assert.Equal(t, c.wantCode, rr.Code)
		})
	}
}

func TestSignUpErr400(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		req, err := http.NewRequest("POST", conf.PathSignUp, strings.NewReader("ASDSAD"))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		userHandler := NewUserHandler()
		userHandler.SignUp(rr, req)
		assert.Equal(t, 400, rr.Code)
	})
}
