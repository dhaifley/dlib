package dlib

import (
	"database/sql"
)

// SQLRows types represent cursors iterating over SQL query results.
type SQLRows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

// SQLExecutor types are able to execute SQL statements.
type SQLExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (SQLRows, error)
	Close() error
	Ping() error
	Stats() sql.DBStats
}

// SQLSession values wrap the sql database driver for layered abatraction
type SQLSession struct {
	DB *sql.DB
}

// Exec abstracts the sql database driver exec function
func (ss *SQLSession) Exec(query string, args ...interface{}) (sql.Result, error) {
	return ss.DB.Exec(query, args...)
}

// Query abstracts the sql database driver query function
func (ss *SQLSession) Query(query string, args ...interface{}) (SQLRows, error) {
	return ss.DB.Query(query, args...)
}

// Close abstracts the sql database driver close function
func (ss *SQLSession) Close() error {
	if ss.DB != nil {
		return ss.DB.Close()
	}

	return nil
}

// Ping abstracts the sql database driver ping function
func (ss *SQLSession) Ping() error {
	return ss.DB.Ping()
}

//Stats abstracts the sql database driver stats function
func (ss *SQLSession) Stats() sql.DBStats {
	return ss.DB.Stats()
}
