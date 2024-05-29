package tests

import (
	"testing"

	"github.com/LuizGuilherme13/clock-in-api/db"
)

func TestDBConn(t *testing.T) {
	if _, err := db.OpenConn(); err != nil {
		t.Fatal(err)
	}
}
