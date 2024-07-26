package altcha

import (
	"github.com/tundrawork/powcha/config"
	"testing"
)

const DefaultComplexity = 100000

func TestParametersPopulate(t *testing.T) {
	config.Conf.Altcha = config.Altcha{
		Algorithm:  "SHA-256",
		Complexity: 100000,
	}

	params := Parameters{}

	params.Populate()

	if params.Algorithm != "SHA-256" {
		t.Errorf("Expected default algorithm SHA-256, got %s", params.Algorithm)
	}

	if len(params.Salt) < 10 {
		t.Errorf("Expected salt to be at least 10 characters, got %s", params.Salt)
	}

	if params.Complexity != DefaultComplexity {
		t.Errorf("Expected default complexity %d, got %d", DefaultComplexity, params.Complexity)
	}

	if params.Number <= 0 || params.Number < MinimumComplexity || params.Number > params.Complexity {
		t.Errorf("Number is not in the expected range: got %d", params.Number)
	}
}
