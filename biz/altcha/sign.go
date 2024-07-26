package altcha

import (
	"crypto/hmac"
	"encoding/base64"
	"github.com/tundrawork/powcha/biz/altcha/internal"
	"hash"

	"github.com/tundrawork/powcha/biz/altcha/algorithm"
)

// Sign generates a signature for the given text.
func Sign(algo algorithm.Algorithm, text string) string {
	secret, _ := GetSecrets()
	return sign(algo, text, secret)
}

func sign(algo algorithm.Algorithm, text, secret string) string {
	if len(secret) == 0 {
		panic("secret not provided to signing function")
	}
	newHasher := func() (hasher hash.Hash) {
		hasher, _ = internal.GetHasher(algo)
		return hasher
	}
	signer := hmac.New(newHasher, []byte(secret))
	signer.Write([]byte(text))

	// The official server implementation example uses hex encoding.
	// Change to this doesn't affect compatibility since the client doesn't read the signature.
	// This implementation uses base64 encoding produces shorter signatures.
	return base64.RawURLEncoding.EncodeToString(signer.Sum(nil))
}

// VerifySignature checks if the given signature is valid for the given text.
func VerifySignature(algo algorithm.Algorithm, text string, signature string) (valid bool) {
	if len(signature) == 0 {
		return false
	}

	current, previous := GetSecrets()

	// Check using the current secret
	validSignature := sign(algo, text, current)
	if signature == validSignature {
		return true
	}

	// Check using the previous secret
	validSignature = sign(algo, text, previous)
	if signature == validSignature {
		return true
	}

	return false
}
