package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	InitDB()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestYourFunc(t *testing.T) {
	// Test code
}
