package data

import (
	"github.com/bits-and-blooms/bloom"
)

type Filter interface {
	CheckUsername(string) bool
	IsAlive() bool
	AddEntry(string)
}

type db struct {
	*bloom.BloomFilter
}

func NewDB() Filter {
	return &db{
		bloom.NewWithEstimates(100000, 0.01),
	}
}

// // Test
// func init() {
// 	filter.Add([]byte("nassah221"))
// }

func (d *db) CheckUsername(username string) bool {
	return d.Test([]byte(username))
}

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

// For testing
func (d *db) AddEntry(s string) {
	d.Add([]byte(s))
}
