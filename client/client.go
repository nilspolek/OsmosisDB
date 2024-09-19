package client

import (
	"bufio"
	"errors"
	"net"

	"github.com/nilspolek/OsmosisDB/paser"
)

type OsmosisDB struct {
	conn net.Conn
}

func NewOsmosisDB(Address string) (*OsmosisDB, error) {
	conn, err := net.Dial("tcp", Address)
	if err != nil {
		return nil, err
	}
	return &OsmosisDB{conn}, nil
}

func (o *OsmosisDB) Set(Key string, Value []byte, dataType byte) error {
	command := paser.Command{
		Type:     paser.SET,
		Keyword:  Key,
		DataType: dataType,
		Data:     Value,
	}
	_, err := o.conn.Write(command.Bytes())
	if err != nil {
		return err
	}
	buffer := bufio.NewReader(o.conn)
	response, err := buffer.ReadBytes(byte('\n'))
	if err != nil {
		return err
	}
	pser := paser.NewPaser()
	responseCommand, err := pser.Parse(response)
	if responseCommand.Type == paser.ERR {
		return errors.New(responseCommand.Keyword)
	}
	return nil
}

func (o *OsmosisDB) Get(Key string) ([]byte, error) {
	command := paser.Command{
		Type:    paser.GET,
		Keyword: Key,
	}
	_, err := o.conn.Write(command.Bytes())
	if err != nil {
		return nil, err
	}
	buffer := bufio.NewReader(o.conn)
	response, err := buffer.ReadBytes(byte('\n'))
	pser := paser.NewPaser()
	responseCommand, err := pser.Parse(response)
	if responseCommand.Type == paser.ERR {
		return nil, errors.New(responseCommand.Keyword)
	}
	return []byte(responseCommand.Keyword), err
}

func (o *OsmosisDB) Delete(Key string) error {
	command := paser.Command{
		Type:    paser.DEL,
		Keyword: Key,
	}
	_, err := o.conn.Write(command.Bytes())
	if err != nil {
		return err
	}
	buffer := bufio.NewReader(o.conn)
	response, err := buffer.ReadBytes(byte('\n'))
	pser := paser.NewPaser()
	responseCommand, err := pser.Parse(response)
	if responseCommand.Type == paser.ERR {
		return errors.New(responseCommand.Keyword)
	}
	return nil
}

func (o *OsmosisDB) Update(Key string, Value []byte, dataType byte) error {
	command := paser.Command{
		Type:     paser.UPT,
		Keyword:  Key,
		DataType: dataType,
		Data:     Value,
	}
	o.conn.Write(command.Bytes())
	buffer := bufio.NewReader(o.conn)
	response, err := buffer.ReadBytes(byte('\n'))
	if err != nil {
		return err
	}
	pser := paser.NewPaser()
	responseCommand, err := pser.Parse(response)
	if responseCommand.Type == paser.ERR {
		return errors.New(responseCommand.Keyword)
	}
	return nil
}

func (o *OsmosisDB) Close() error {
	return o.conn.Close()
}
