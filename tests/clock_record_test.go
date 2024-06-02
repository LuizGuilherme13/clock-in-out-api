package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/clock"
	"github.com/LuizGuilherme13/clock-in-api/employee"
)

func TestClockRecord(t *testing.T) {
	mux := http.NewServeMux()
	clock.Controller(mux)
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("post clock", func(t *testing.T) {

		employee := employee.Model{
			UserName: "luiz",
			Password: "12345",
		}

		body, err := json.Marshal(employee)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post(server.URL+"/clock", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("get clock", func(t *testing.T) {
		eeId := 5

		resp, err := http.Get(fmt.Sprintf("%s/clock/%d", server.URL, eeId))
		if err != nil {
			t.Fatal(err)
		}

		clocks := clock.Records{}
		if err := json.NewDecoder(resp.Body).Decode(&clocks); err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", resp.StatusCode, http.StatusOK)
		}
	})
}
