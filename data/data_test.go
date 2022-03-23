package data

import (
	"flare/exercise/utils"
	"testing"

	"github.com/bits-and-blooms/bloom"
)

func TestBloomField(t *testing.T) {
	f := &db{
		bloom.NewWithEstimates(100000, 0.01),
	}

	var validNames, invalidNames []string

	// Generate some names and insert into the filter
	for i := 0; i <= 100; i++ {
		n := utils.RandomInt(8, 12)
		name := utils.RandomStringValid(int(n))
		validNames = append(validNames, name)

		f.Add([]byte(name))
	}

	// Generate some name and don't insert into the filter
	for i := 0; i <= 100; i++ {
		n := utils.RandomInt(8, 12)
		name := utils.RandomStringValid(int(n))

		invalidNames = append(invalidNames, name)
	}

	// Check if the inserted names do exist
	for _, name := range validNames {
		ok := f.Test([]byte(name))
		if !ok {
			t.Errorf("Name: %s was inserted but not found in filter", name)
		}
	}

	// Check if the not inserted names do show up - which is rather stupid to do but test anyway
	for _, name := range invalidNames {
		ok := f.Test([]byte(name))
		if ok {
			t.Errorf("Name: %s was not inserted but found in filter", name)
		}
	}
}
