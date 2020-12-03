package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestGetAllTimezones(t *testing.T) {
	tzones, err := GetAllTimezones()
	if err != nil {
		t.Error("Get All Timezones failed.")
	}

	if len(tzones) == 0 {
		t.Error("Timezones did not return any values.")
	}
}
