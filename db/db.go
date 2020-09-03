// Package db exposes a very simple database with different implementations
package db

// Db defines the base interface for using de database. Most of the methods return a bool that indicates the success of the operation.
type Db interface {

	// Connect creates the database connection
	Connect() bool

	// Connected returns the connection status
	Connected() bool

	// Disconnect closes the database connection
	Disconnect()

	// Create adds a new object to the database and returns the id of the newly created object
	Create(interface{}) (string, bool)

	// Retrieve returns an object in the database by it's id
	Retrieve(string) (interface{}, bool)

	// Update replaces an object in the database identified with the provided id with another object
	Update(string, interface{}) bool

	// Delete removes and object from the database by it's id
	Delete(string) bool
}
