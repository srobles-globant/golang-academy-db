/*
Package db exposes a very simple database with different implementations
*/
package db

import (
	"fmt"
)

/*
InMemoryDb is the implementation of the Db interface in a map[string]interface{} type
*/
type InMemoryDb struct {
	store            map[string]interface{}
	connected        bool
	idGeneratorCount int
}

// Connect creates the database connection
func (imd *InMemoryDb) Connect() bool {
	if imd.connected {
		return true
	}
	imd.store = make(map[string]interface{})
	imd.connected = true
	imd.idGeneratorCount = 0
	return true
}

// Disconnect closes the database connection
func (imd *InMemoryDb) Disconnect() {
	imd.connected = false
	return
}

// Create adds a new object to the database and returns the id of the newly created object
func (imd *InMemoryDb) Create(obj interface{}) (string, bool) {
	if !imd.connected {
		return "", false
	}
	id := fmt.Sprintf("obj%d", imd.idGeneratorCount)
	imd.idGeneratorCount++
	imd.store[id] = obj
	return id, true
}

// Retrieve returns an object in the database by it's id
func (imd *InMemoryDb) Retrieve(string) (interface{}, bool) {
	return nil, false
}

// Update replaces an object in the database identified with the provided id with another object
func (imd *InMemoryDb) Update(string, interface{}) bool {
	return false
}

// Delete removes and object from the database by it's id
func (imd *InMemoryDb) Delete(string) bool {
	return false
}
