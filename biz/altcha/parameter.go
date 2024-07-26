package altcha

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/tundrawork/powcha/biz/altcha/algorithm"
	"github.com/tundrawork/powcha/biz/altcha/internal"
	"github.com/tundrawork/powcha/config"
)

// MinimumComplexity is the minimum complexity allowed.
// @see https://altcha.org/docs/complexity
const MinimumComplexity = 10000

// Parameters are the parameters used to generate a challenge. If any of the
// parameters are missing, they will be generated.
type Parameters struct {
	// Algorithm is the hashing algorithm used to generate the challenge.
	// Supported algorithms are SHA-256, SHA-384, and SHA-512.
	Algorithm string `json:"algorithm"`

	// Salt is a random string used to generate the challenge.
	// The minimum length is 10 characters.
	Salt string `json:"salt"`

	// Complexity is the number of iterations used to generate the challenge.
	// This is only considered when Number is not provided.
	Complexity int `json:"complexity,omitempty"`

	// Number is the secret number which the client must solve for.
	Number int `json:"number,omitempty"`
}

// Populate generates any missing parameters.
func (params *Parameters) Populate() {
	// Use algorithm from config if not provided.
	algo, ok := algorithm.AlgorithmFromString(params.Algorithm)
	if !ok {
		algo, ok = algorithm.AlgorithmFromString(config.Conf.Altcha.Algorithm)
		if !ok {
			logger.Infof("invalid algorithm in config: %s, fallback to SHA-256", config.Conf.Altcha.Algorithm)
			algo = algorithm.SHA256
		}
	}
	params.Algorithm = algo.String()

	// Without salt, we generate a new one.
	if len(params.Salt) < 16 {
		params.Salt = internal.RandomString(16)
	}

	// Generate random number if not provided.
	// Use complexity from config if not provided.
	if params.Number <= 0 {
		if params.Complexity <= MinimumComplexity {
			params.Complexity = config.Conf.Altcha.Complexity
		}
		params.Number = internal.RandomInt(MinimumComplexity, params.Complexity)
	}
}
