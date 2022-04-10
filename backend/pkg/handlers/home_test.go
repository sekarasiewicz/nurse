package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/sekarasiewicz/nurse/backend/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	t.Run("can get home", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(HomeHandler)

		handler.ServeHTTP(recorder, req)

		// Assert with lib
		assert.Equal(t, recorder.Code, http.StatusOK, "handler returned wrong status")
		// Test status
		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Test string body
		expected := `{"id":1,"name":"test 2"}`

		// have to remove \n at the end of the file
		if strings.TrimRight(recorder.Body.String(), "\n") != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Body.String(), expected)
		}

		// Test struct body
		var user models.User
		err = json.NewDecoder(recorder.Body).Decode(&user)
		if err != nil {
			t.Error("error parsing body")
		}

		expectedUser := &models.User{ID: 1, Name: "test 2"}

		if reflect.DeepEqual(expectedUser, user) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Body.String(), expected)
		}
	})
}
