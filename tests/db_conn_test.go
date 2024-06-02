package tests

import (
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/database"
)

func TestDBConn(t *testing.T) {
	if _, err := database.OpenConn(); err != nil {
		t.Fatal(err)
	}
}
