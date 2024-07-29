package altcha

import (
	"testing"

	"github.com/tundrawork/powcha/biz/altcha/internal"
)

func TestValidateChallenge(t *testing.T) {

	// Override randomString for deterministic behavior
	internal.RandomString = func(length int) string {
		const fakeRandomString = "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ"
		return fakeRandomString
	}
	// Rotate secrets twice so that the fake randomness is used for both secrets
	RotateSecrets()
	RotateSecrets()

	// Test with a valid encoded message
	validMsg := Message{
		Algorithm: "SHA-256",
		Salt:      "0V5xzYiSFmY1swbb",
		Number:    49500,
		Challenge: "69df4e03d8fffc1d66aeba60384ad28d70caed4bcf10c69f80e0a16666eae6a7",
		Signature: "-gytD6e0qjPZknud02kOzq8KqsayfXfGI1exZXFjI6k",
	}
	ok, err := ValidateResponse(validMsg.EncodeWithBase64(), false)
	if err != nil || !ok {
		t.Error("Expected valid encoded message to return true, got false")
	}

	// Test with invalid encoded string
	invalidEncoded := "invalid-base64"
	ok, err = ValidateResponse(invalidEncoded, false)
	if err == nil || ok {
		t.Error("Expected invalid encoded string to return false, got true")
	}

	// Test with a message which has valid encoding but invalid values

	// Test with an invalid challenge
	invalidChallengeMsg := validMsg
	invalidChallengeMsg.Challenge = "incorrect_challenge"
	ok, err = ValidateResponse(invalidChallengeMsg.EncodeWithBase64(), false)
	if err != nil || ok {
		t.Error("Expected message with invalid challenge to be false, got true")
	}

	// Test with invalid signature
	invalidSignatureMsg := validMsg
	invalidSignatureMsg.Signature = "incorrect_signature"
	ok, err = ValidateResponse(invalidSignatureMsg.EncodeWithBase64(), false)
	if err != nil || ok {
		t.Error("Expected message with invalid signature to be false, got true")
	}

}

func TestValidateChallengeReplayPrevention(t *testing.T) {

	// Override randomString for deterministic behavior
	internal.RandomString = func(length int) string {
		const fakeRandomString = "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ"
		return fakeRandomString
	}
	// Rotate secrets twice so that the fake randomness is used for both secrets
	RotateSecrets()
	RotateSecrets()

	// Generate a valid encoded response
	validMsg := Message{
		Algorithm: "SHA-256",
		Salt:      "0V5xzYiSFmY1swbb",
		Number:    49500,
		Challenge: "69df4e03d8fffc1d66aeba60384ad28d70caed4bcf10c69f80e0a16666eae6a7",
		Signature: "-gytD6e0qjPZknud02kOzq8KqsayfXfGI1exZXFjI6k",
	}
	response := validMsg.EncodeWithBase64()

	// The first time should be valid
	ok, err := ValidateResponse(response, true)
	if err != nil || !ok {
		t.Error("Expected first use of response to return true, got false")
	}

	// The second time should be invalid
	ok, err = ValidateResponse(response, true)
	if err != nil || ok {
		t.Error("Expected second use of response to return false, got true")
	}

}
