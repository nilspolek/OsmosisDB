package main

import (
	"github.com/nilspolek/OsmosisDB/server"
	"github.com/nilspolek/goLog"
)

func main() {
	server := server.NewServer(server.NewServerConfig(":8080"))
	goLog.Info("Starting server on port 8080")
	server.Start()
	defer server.Stop()

}
