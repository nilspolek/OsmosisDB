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

	command := new(Command)

	var err error
	command.Type, err = checkCommand(data[:3])
	if err != nil {
		return *command, err
	}

	//payload := strings.Split(string(data[2:]), "\r\n") //"SETNAME\r\n@NILS\n"

	idx := bytes.Index(data, []byte("\r\n"))

	command.Keyword = string(data[3:idx])

	command.DataType = data[idx+2]

	command.Data = data[idx+3:]

	return *command, nil
}

func checkCommand(test []byte) (string, error) { //(p *Paser) ?

	switch string(test) {
	case SET:
		return SET, nil
	case GET:
		return GET, nil
	case DEL:
		return DEL, nil
	case UPT:
		return UPT, nil
	case ERR:
		return ERR, nil
	case OK:
		return OK, nil
	default:
		return "", errors.New("Command type not found")
	}

}
