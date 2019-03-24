package cupsgenerator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	got := Generate()

	if len(got) != 20 {
		t.Errorf("Wrong lenght")
	}

}

