package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/nilspolek/OsmosisDB/paser"
)

// Service is a simple key-value store
type Service struct {
	dbFileName string
	data       map[string][]byte
}

// NewService creates a new database service
func NewService(filename string) (*Service, error) {
	data, err := loadMapFromFile(filename)
	if err != nil {
		data = make(map[string][]byte)
	}
	db := Service{
		dbFileName: filename,
		data:       data,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(db *Service) {
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

// Command executes a command on the database
func (d *Service) Command(command paser.Command) ([]byte, error) {
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

// Get returns the value for a key or an error if the key does not exist
func (d *Service) Get(key string) ([]byte, error) {
	if d.data[key] == nil {
		return nil, fmt.Errorf("key [%s] not found", key)
	}
	return d.data[key], nil
}

// Set sets the value for a key or returns an error if the key already exists
func (d *Service) Set(key string, value []byte) error {
	if d.data[key] == nil {
		d.data[key] = value
		return nil
	}
	return fmt.Errorf("key already exists %s", key)
}

// Delete deletes a key from the database
func (d *Service) Delete(key string) {
	delete(d.data, key)
}

// Update updates the value for a key or returns an error if the key does not exist
func (d *Service) Update(key string, value []byte) error {
	if _, ok := d.data[key]; !ok {
		return fmt.Errorf("key [%s] not found", key)
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

// Close saves the database to a file
func (d *Service) Close() error {
	return saveMapToFile(d.data, d.dbFileName)
}
