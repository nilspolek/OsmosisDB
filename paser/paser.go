package paser

import (
	"bytes"
	"errors"
)

type Paser struct{}

type Command struct {
	Type     string
	Keyword  string
	DataType byte
	Data     []byte
}

const (
	TYPEINT    = '!'
	TYPESTRING = '@'
	TYPEBOOL   = '#'
	TYPEFLOAT  = '$'
	SET        = "SET"
	GET        = "GET"
	DEL        = "DEL"
	UPT        = "UPT"
	ERR        = "ERR"
	OK         = "OK "
)

func NewPaser() *Paser {
	return &Paser{}
}

func (p *Paser) Parse(data []byte) (Command, error) {
	// Parse the input
	var (
		err     error
		command Command
	)
	switch string(data[:3]) {
	case SET:
		command, err = p.parseSet(data)
		break
	case GET:
		command, err = p.parseGet(data)
		break
	case DEL:
		command, err = p.parseDel(data)
		break
	case UPT:
		command, err = p.parseUpt(data)
		break
	case ERR:
		command, err = p.parseErr(data)
		break
	case OK:
		command, err = Command{Type: OK}, nil
		break
	default:
		err = errors.New("Command type not found")
		break
	}
	if err != nil {
		return Command{}, err
	}
	return command, nil
}

func (p *Paser) parseErr(data []byte) (Command, error) {
	var (
		command Command
	)
	command.Type = ERR
	command.Keyword = string(data[3 : len(data)-1])
	return command, nil
}

func (p *Paser) parseUpt(data []byte) (Command, error) {
	var (
		command Command
		idx     int
	)
	idx = bytes.Index(data, []byte("\r\n"))
	command.Type = UPT
	command.Keyword = string(data[3:idx])
	command.DataType = data[idx+2]
	command.Data = data[idx+3 : len(data)-1]
	return command, nil
}

func (p *Paser) parseDel(data []byte) (Command, error) {
	var (
		command Command
	)
	command.Type = DEL
	command.Keyword = string(data[3 : len(data)-1])
	return command, nil
}

func (p *Paser) parseSet(data []byte) (Command, error) {
	var (
		command Command
		idx     int
	)
	idx = bytes.Index(data, []byte("\r\n"))
	command.Type = SET
	command.Keyword = string(data[3:idx])
	command.DataType = data[idx+2]
	command.Data = data[idx+3 : len(data)-1]
	return command, nil
}
func (p *Paser) parseGet(data []byte) (Command, error) {
	var (
		command Command
	)
	command.Type = GET
	command.Keyword = string(data[3 : len(data)-1])
	return command, nil
}
