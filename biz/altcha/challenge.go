package altcha

import (
	"github.com/tundrawork/powcha/biz/altcha/algorithm"
	"github.com/tundrawork/powcha/biz/altcha/internal"
)

// NewChallenge creates a new challenge with default parameters.
func NewChallenge() (msg Message) {
	return NewChallengeWithParams(Parameters{})
}

// NewChallengeEncoded creates a new challenge with default parameters and encoded for the client.
func NewChallengeEncoded() string {
	// Create a new challenge message.
	msg := NewChallengeWithParams(Parameters{})

	// Return the encoded challenge message.
	return msg.Encode()
}

// NewChallengeWithParams creates a new challenge with the given parameters.
func NewChallengeWithParams(params Parameters) (msg Message) {
	// Populate any missing parameters.
	params.Populate()

	// Generate the challenge and signature.
	algo, _ := algorithm.AlgorithmFromString(params.Algorithm)
	challenge := internal.GenerateHash(algo, params.Salt, params.Number)
	signature := Sign(algo, challenge)
	msg = Message{
		Algorithm: params.Algorithm,
		Salt:      params.Salt,
		Challenge: challenge,
		Signature: signature,
		// Number is a secret and must not be exposed to the client.
	}

	// Return the challenge message.
	return msg
}
