package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/nilspolek/OsmosisDB/paser"
)

type Server struct {
	ln     net.Listener
	config ServerConfig
}
type ServerConfig struct {
	addr string
}

func NewServerConfig(addr string) ServerConfig {
	return ServerConfig{
		addr: addr,
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
		go handleConnection(conn)
	}
}

func (s *Server) Stop() error {
	return s.ln.Close()
}

func handleConnection(conn net.Conn) {
	var (
		command paser.Command
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
		if err != nil {
			conn.Write(paser.Command{Type: paser.ERR, Keyword: err.Error()}.Bytes())
			continue
		}

		conn.Write([]byte([]byte(fmt.Sprintf("Command Type:\t%s\nCommand Keyword:\t%s\nCommand Datatype\t%v\nCommand Data\t%s\n", command.Type, command.Keyword, string(command.DataType), string(command.Data)))))
	}
}
