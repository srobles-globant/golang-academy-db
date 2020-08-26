/*
Package db exposes a very simple database with different implementations
*/
package db

/*
Db defines the base interface for using de database
*/
type Db interface {
	Connect() bool
	Disconnect()
}
