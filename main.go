package main

import (
	"github.com/nilspolek/OsmosisDB/server"
)

func main() {
	server := server.NewServer(server.NewServerConfig(":8080"))
	server.Start()
	defer server.Stop()

}
