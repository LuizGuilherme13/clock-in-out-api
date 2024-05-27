package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/controllers"
)

func TestEmployee(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/employee", nil)
		rr := httptest.NewRecorder()

		handle := http.HandlerFunc(controllers.EmployeeController)
		handle.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf("rr.Code = %d, expected = %d", rr.Code, http.StatusCreated)
		}
	})
}
