package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sekarasiewicz/nurse/backend/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	t.Run("can get home", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/", HomeHandler)
		router.ServeHTTP(w, req)

		// Assert with lib
		assert.Equal(t, w.Code, http.StatusOK, "handler returned wrong status")
		// Test status
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Test struct body
		var user models.User
		err = json.NewDecoder(w.Body).Decode(&user)
		if err != nil {
			t.Error("error parsing body")
		}

		expectedUser := &models.User{ID: 1, Name: "test 2"}

		if reflect.DeepEqual(expectedUser, user) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expectedUser)
		}
	})
}
