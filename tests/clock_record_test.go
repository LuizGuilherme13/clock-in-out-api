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
	t.Run("post punch", func(t *testing.T) {

		employee := employee.Model{
			UserName: "luiz",
			Password: "12345",
		}

		body, err := json.Marshal(employee)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/punch", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handle := http.HandlerFunc(clock.Controller)
		handle.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", rr.Code, http.StatusOK)
		}
	})

	t.Run("get punches", func(t *testing.T) {
		eeId := 5

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/punch?id=%d", eeId), nil)
		rr := httptest.NewRecorder()

		handle := http.HandlerFunc(clock.Controller)
		handle.ServeHTTP(rr, req)

		clocks := clock.Punches{}
		if err := json.NewDecoder(rr.Body).Decode(&clocks); err != nil {
			t.Fatal(err)
		}

		if rr.Code != http.StatusOK {
			t.Fatalf("rr.Code = %d, expected = %d", rr.Code, http.StatusOK)
		}
	})
}
