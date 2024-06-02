package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/employee"
)

func TestEmployee(t *testing.T) {
	mux := http.NewServeMux()
	employee.Controller(mux)
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("create employee", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/employee", "application/json", nil)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", resp.StatusCode, http.StatusOK)
		}
	})
}
