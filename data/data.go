package data

import (
	"github.com/bits-and-blooms/bloom"
)

type DB interface {
	CheckUsername(string) bool
	IsAlive() bool
	AddEntry(string)
}

type db struct {
	*bloom.BloomFilter
}

func NewDB() DB {
	return &db{
		bloom.NewWithEstimates(100000, 0.01),
	}
}

// CheckUsername checks the given username against the db
// returns true if it is found and false otherwise
func (d *db) CheckUsername(username string) bool {
	return d.Test([]byte(username))
}

// IsAlive returns true if the service is available and false otherwise
// Normally this function would ping the service database connection
func (d *db) IsAlive() bool {
	// Normally implementing a health check would start with pinging
	// any external services that our service might rely on
	// however, for the scope of this exercise that is trivial
	//
	// Therefore, I'm doing the following check which is obviously naive
	// but is ok for this exercise

	// Check if our bloom filter has been initialized
	return d != nil
}

// AddEntry adds usernames to the in-memory db
// so that it is much simpler to test with auto-generated data
func (d *db) AddEntry(s string) {
	d.Add([]byte(s))
}
