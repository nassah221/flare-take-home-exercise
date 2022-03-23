package utils

import (
	"strings"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("Test invalid username string", func(t *testing.T) {
		for i := 0; i <= 100; i++ {
			n := RandomInt(1, 20)
			name := RandomStringInvalid(int(n))

			if int(n) != len(name) {
				t.Fatalf("Name %s is of length %d instead of %d", name, len(name), n)
			}

			if len(name) >= 8 && len(name) <= 12 {
				if !strings.ContainsAny(string(name[0]), numbers) {
					t.Logf("%v", len(name) >= 8 && len(name) <= 12)
					t.Fatalf("Name [%s] failed test", name)
				}
			}
		}
	})
}
