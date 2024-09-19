package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/nilspolek/OsmosisDB/database"
	"github.com/nilspolek/OsmosisDB/paser"
)

type Server struct {
	ln     net.Listener
	config ServerConfig
}
type ServerConfig struct {
	addr string
	db   database.DatabaseService
}

func NewServerConfig(addr string, db *database.DatabaseService) ServerConfig {
	return ServerConfig{
		addr: addr,
		db:   *db,
	}
}

func NewServer(config ServerConfig) *Server {
	return &Server{
		config: config,
	}
}

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

func (s *Server) Stop() error {
	return s.ln.Close()
}

func (s *Server) handleConnection(conn net.Conn) {
	var (
		command paser.Command
		result  []byte
	)
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())
	conn.Write([]byte("OK \n"))
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
				Type:     paser.ERR,
				Keyword:  err.Error(),

			}.Bytes())
			continue
		}
		conn.Write(paser.Command{
			Type:     paser.OK,
			Keyword:  string(result),
		}.Bytes())
	}
}
