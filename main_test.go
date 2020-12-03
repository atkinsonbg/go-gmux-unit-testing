package main

import (
	"os"
	"testing"

	database "github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	database.InitDB()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}
