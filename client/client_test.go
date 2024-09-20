package client

import (
	"os"
	"testing"

	"github.com/nilspolek/OsmosisDB/database"
	"github.com/nilspolek/OsmosisDB/server"
	"github.com/nilspolek/goLog"
)

var (
	db  *database.Service
	err error
)

// Sets up the env for the tests
func TestMain(m *testing.M) {
	go func() {
		db, err = database.NewService("db.json")
		if err != nil {
			goLog.Error(err.Error())
			os.Exit(1)
		}
		defer db.Close()
		server := server.NewServer(server.NewConfig(":8080", db))
		server.Start()
		defer server.Stop()
	}()
}

func TestSet(t *testing.T) {
	client, _ := NewOsmosisDB("localhost:8080")
	defer client.Close()
	err := client.Set("test", []byte("test123"), byte('@'))
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestGet(t *testing.T) {
	client, err := NewOsmosisDB("localhost:8080")
	defer client.Close()
	if err != nil {
		t.Fatal(err)
	}
	val, err := client.Get("test")
	if err != nil {
		t.Fatal(err)
	}
	if string(val) != "test123" {
		t.Fatalf("Expected 'test123' got %s", string(val))
	}
	db.Close()
}
