package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/nilspolek/OsmosisDB/paser"
)

type DatabaseService struct {
	dbFileName string
	data       map[string][]byte
}

func NewDatabaseService(filename string) (*DatabaseService, error) {

	data, err := loadMapFromFile(filename)
	if err != nil {
		data = make(map[string][]byte)
	}
	db := DatabaseService{
		dbFileName: filename,
		data:       data,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(db *DatabaseService) {
		for sig := range c {
			if sig == os.Interrupt {
				db.Close()
				os.Exit(0)
			}
			if sig == os.Kill {
				db.Close()
				os.Exit(0)
			}
		}
	}(&db)

	return &db, nil
}

func (d *DatabaseService) Command(command paser.Command) ([]byte, error) {
	var (
		result []byte
		err    error
	)
	switch command.Type {
	case paser.GET:
		result, err = d.Get(command.Keyword)
	case paser.SET:
		err = d.Set(command.Keyword, command.Data)
	case paser.DEL:
		d.Delete(command.Keyword)
	case paser.UPT:
		err = d.Update(command.Keyword, command.Data)
	default:
		err = errors.New("Command type not found")
	}
	return result, err
}

func (d *DatabaseService) Get(key string) ([]byte, error) {
	if d.data[key[:len(key)-1]] == nil {
		return nil, errors.New(fmt.Sprintf("key [%s] not found", key))
	}

	return d.data[key[:len(key)-1]], nil
}

func (d *DatabaseService) Set(key string, value []byte) error {
	if d.data[key] == nil {
		d.data[key] = value
		return nil
	}
	return errors.New(fmt.Sprintf("key already exists %s", key))
}

func (d *DatabaseService) Delete(key string) {
	delete(d.data, key)
}

func (d *DatabaseService) Update(key string, value []byte) error {
	if _, ok := d.data[key]; !ok {
		return errors.New(fmt.Sprintf("key [%s] not found", key))
	}
	d.data[key] = value
	return nil
}

func saveMapToFile(data map[string][]byte, filename string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Speichere die JSON-Daten in die Datei
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadMapFromFile(filename string) (map[string][]byte, error) {
	var data map[string][]byte

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		os.Create(filename)
		err := os.WriteFile(filename, []byte("{}"), 0644)
		if err != nil {
			return nil, err
		}
	}
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Konvertiere JSON zurück in die Map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DatabaseService) Close() error {
	return saveMapToFile(d.data, d.dbFileName)
}
