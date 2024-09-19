package main

import (
	"os"

	"github.com/nilspolek/OsmosisDB/database"
	"github.com/nilspolek/OsmosisDB/server"
	"github.com/nilspolek/goLog"
)

func main() {
	db, err := database.NewDatabaseService("db.json")
	if err != nil {
		goLog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	server := server.NewServer(server.NewServerConfig(":8080", db))
	server.Start()
	defer server.Stop()
}
