package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/nilspolek/OsmosisDB/database"
	"github.com/nilspolek/OsmosisDB/paser"
)

// Server struct for the server that handles the connections
type Server struct {
	ln     net.Listener
	config Config
}

// Config struct for the server configuration
type Config struct {
	addr string
	db   database.Service
}

// NewConfig create a new server configuration
func NewConfig(addr string, db *database.Service) Config {
	return Config{
		addr: addr,
		db:   *db,
	}
}

// NewServer create a new server
func NewServer(config Config) *Server {
	return &Server{
		config: config,
	}
}

// Start the server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.config.addr)
	if err != nil {
		return err
	}
	s.ln = ln
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// Stop the server
func (s *Server) Stop() error {
	s.config.db.Close()
	return s.ln.Close()
}

func (s *Server) handleConnection(conn net.Conn) {
	var (
		command paser.Command
		result  []byte
	)
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())
	buffer := bufio.NewReader(conn)
	pser := paser.NewPaser()
	for {
		clientInput, err := buffer.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}
		command, err = pser.Parse([]byte(clientInput))
		result, err = s.config.db.Command(command)
		if err != nil {
			conn.Write(paser.Command{
				Type:    paser.ERR,
				Keyword: err.Error(),
			}.Bytes())
			continue
		}

		conn.Write(paser.Command{
			Type:    paser.OK,
			Keyword: string(result),
		}.Bytes())
	}
}
