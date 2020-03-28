package cupsgenerator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	for i := 0; i < 10000; i++ {
		got := Generate()

		if len(got) != 20 {
			t.Errorf("Wrong lenght")
		}
	}
}
