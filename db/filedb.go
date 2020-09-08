// Package db exposes a very simple database with different implementations
package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FileDb is the implementation of the Db interface in a map[string]interface{} type and file persistance
type FileDb struct {
	FilePath string

	filePathFixed string
	InMemoryDb
}

// Connect creates the database connection
func (fd *FileDb) Connect() bool {
	if fd.connected {
		return false
	}

	file, err := os.OpenFile(fd.FilePath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return false
	}
	fd.filePathFixed = fd.FilePath
	defer file.Close()

	var data []byte
	data, err = ioutil.ReadAll(file)
	if len(data) > 0 && !json.Valid(data) {
		return false
	}

	if len(data) == 0 {
		fd.idGenCount = 0
		fd.store = make(map[string]interface{})
	} else {
		json.Unmarshal(data, &fd.store)
		fd.idGenCount = int(fd.store["count"].(float64))
	}

	fd.connected = true
	return true
}

// Disconnect closes the database connection
func (fd *FileDb) Disconnect() {
	defer func() {
		fd.connected = false
	}()

	fd.store["count"] = fd.idGenCount

	file, err := os.OpenFile(fd.filePathFixed, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := json.Marshal(fd.store)
	if err != nil {
		return
	}

	file.Write(data)
}
