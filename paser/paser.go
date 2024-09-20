package paser

import (
	"bytes"
	"errors"
)

// Paser struct
type Paser struct{}

// Command struct
type Command struct {
	Type     string
	Keyword  string
	DataType byte
	Data     []byte
}

const (
	// TYPEINT = int used for the data type of the command
	TYPEINT = '!'
	// TYPESTRING = string used for the data type of the command
	TYPESTRING = '@'
	// TYPEBOOL = bool used for the data type of the command
	TYPEBOOL = '#'
	// TYPEFLOAT = float64 used for the data type of the command
	TYPEFLOAT = '$'
	// TYPEBYTE = byte used for the data type of the command
	TYPEBYTE = '%'

	// SET = set command type
	SET = "SET"
	// GET = get command type
	GET = "GET"
	// DEL = delete command type
	DEL = "DEL"
	// UPT = update command type
	UPT = "UPT"
	// ERR = error command type
	ERR = "ERR"
	// OK = ok command type
	OK = "OK "
	// DELIMITER = is the delimiter for the data
	DELIMITER = ";"
)

// NewCommand create a new command
func NewCommand() *Command {
	return &Command{}
}

// Bytes return the command as a byte array
func (c Command) Bytes() []byte {
	return []byte(c.String())
}

// String return the command as a string
func (c Command) String() string {
	if c.Type == SET || c.Type == UPT {
		return c.Type + c.Keyword + DELIMITER + string(c.DataType) + string(c.Data) + "\n"
	}
	return c.Type + c.Keyword + string(c.Data) + "\n"
}

// NewPaser create a new paser
func NewPaser() *Paser {
	return &Paser{}
}

// Parse the input data to a Command
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
		command, err = p.parseOk(data)
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

func (p *Paser) parseOk(data []byte) (Command, error) {
	var (
		command Command
	)
	command.Type = OK
	command.Keyword = string(data[3 : len(data)-1])
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
	idx = bytes.Index(data, []byte(DELIMITER))
	command.Type = UPT
	command.Keyword = string(data[3:idx])
	command.DataType = data[idx+1]
	command.Data = data[idx+2 : len(data)-1]
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
	idx = bytes.Index(data, []byte(DELIMITER))
	command.Type = SET
	command.Keyword = string(data[3:idx])
	command.DataType = data[idx+1]
	command.Data = data[idx+2 : len(data)-1]
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
