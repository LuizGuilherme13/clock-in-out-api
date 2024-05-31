package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/controllers"
	"github.com/LuizGuilherme13/clock-in-api/models"
)

func TestTimeEntries(t *testing.T) {
	t.Run("punch", func(t *testing.T) {

		employee := models.Employee{
			UserName: "luiz",
			Password: "12345",
		}

		body, err := json.Marshal(employee)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/punch", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handle := http.HandlerFunc(controllers.TimeEntriesController)
		handle.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", rr.Code, http.StatusOK)
		}
	})
}
