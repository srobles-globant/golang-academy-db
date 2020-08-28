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
	IDGeneratorPrefix string

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
	if imd.IDGeneratorPrefix == "" {
		imd.IDGeneratorPrefix = "obj"
	}
	return true
}

// Disconnect closes the database connection
func (imd *InMemoryDb) Disconnect() {
	imd.connected = false
	return
}

// Connected returns the connection status
func (imd *InMemoryDb) Connected() bool {
	return imd.connected
}

// Create adds a new object to the database and returns the id of the newly created object
func (imd *InMemoryDb) Create(obj interface{}) (string, bool) {
	if !imd.connected {
		return "", false
	}
	id := fmt.Sprintf("%s%d", imd.IDGeneratorPrefix, imd.idGeneratorCount)
	imd.idGeneratorCount++
	imd.store[id] = obj
	return id, true
}

// Retrieve returns an object in the database by it's id
func (imd *InMemoryDb) Retrieve(id string) (interface{}, bool) {
	if !imd.connected {
		return nil, false
	}
	elem, ok := imd.store[id]
	return elem, ok
}

// Update replaces an object in the database identified with the provided id with another object
func (imd *InMemoryDb) Update(id string, obj interface{}) bool {
	if !imd.connected {
		return false
	}
	if _, ok := imd.store[id]; !ok {
		return false
	}
	imd.store[id] = obj
	return true
}

// Delete removes and object from the database by it's id
func (imd *InMemoryDb) Delete(id string) bool {
	if !imd.connected {
		return false
	}
	if _, ok := imd.store[id]; !ok {
		return false
	}
	delete(imd.store, id)
	return true
}
