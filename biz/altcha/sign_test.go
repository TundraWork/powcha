package altcha

import (
	"github.com/tundrawork/powcha/biz/altcha/algorithm"
	"github.com/tundrawork/powcha/biz/altcha/internal"
	"sync/atomic"
	"testing"
)

func TestSign(t *testing.T) {
	// Override randomString for deterministic behavior
	internal.RandomString = func(length int) string {
		const fakeRandomString = "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ"
		return fakeRandomString
	}
	RotateSecrets() // Rotate secrets so that the fake random string is used

	const exampleText = "The quick brown fox jumps over the lazy dog"
	type args struct {
		algo algorithm.Algorithm
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			algorithm.SHA256.String(),
			args{algorithm.SHA256, exampleText},
			"aHt9C_C2HSU0Bh9bNSQnpmVKPRmZe5BC9Ib6vtil4Lg",
		},
		{
			algorithm.SHA384.String(),
			args{algorithm.SHA384, exampleText},
			"Eisr40pH2bSFBwlk0vMhHgeZXlRusMW8R2jI9qfEdGtC63PdhFdvmHHUnYrNdjNI",
		},
		{
			algorithm.SHA512.String(),
			args{algorithm.SHA512, exampleText},
			"Hvmq3828Zr3AIBnUB6CC50xVB0yNU38srrQAoJY0hLXYwQdy6q56BhbTU36WdmZu-xO2eHKgC2Dex0Rs-hQh5Q",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.algo, tt.args.text); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifySignature(t *testing.T) {
	// Override randomString for deterministic behavior
	internal.RandomString = func(length int) string {
		const fakeRandomString = "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ"
		return fakeRandomString
	}
	RotateSecrets() // Rotate secrets so that the fake random string is used

	const exampleText = "The quick brown fox jumps over the lazy dog"
	type args struct {
		algo      algorithm.Algorithm
		text      string
		signature string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Valid Signature with SHA256",
			args{
				algorithm.SHA256,
				exampleText,
				"aHt9C_C2HSU0Bh9bNSQnpmVKPRmZe5BC9Ib6vtil4Lg",
			},
			true,
		},
		{
			"Invalid Signature with SHA256",
			args{
				algorithm.SHA256,
				exampleText,
				"invalid_signature",
			},
			false,
		},
		{
			"Valid Signature with SHA384",
			args{
				algorithm.SHA384,
				exampleText,
				"Eisr40pH2bSFBwlk0vMhHgeZXlRusMW8R2jI9qfEdGtC63PdhFdvmHHUnYrNdjNI",
			},
			true,
		},
		{
			"Invalid Signature with SHA384",
			args{
				algorithm.SHA384,
				exampleText,
				"invalid_signature",
			},
			false,
		},
		{
			"Valid Signature with SHA512",
			args{
				algorithm.SHA512,
				exampleText,
				"Hvmq3828Zr3AIBnUB6CC50xVB0yNU38srrQAoJY0hLXYwQdy6q56BhbTU36WdmZu-xO2eHKgC2Dex0Rs-hQh5Q",
			},
			true,
		},
		{
			"Invalid Signature with SHA512",
			args{
				algorithm.SHA512,
				exampleText,
				"invalid_signature",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifySignature(tt.args.algo, tt.args.text, tt.args.signature); got != tt.want {
				t.Errorf("VerifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignatureValidityAfterSecretRotations(t *testing.T) {
	// Setup: Override randomString and track calls to generate different secrets
	var rotationCount int32
	internal.RandomString = func(length int) string {
		currentRotation := atomic.LoadInt32(&rotationCount)
		switch currentRotation {
		case 0:
			return "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ" // First secret
		case 1:
			return "1K7xwZjTHfM2tRbbLwJnBgbXcw8zKwXW" // Second secret
		default:
			return "2L8ywAkUIgN3uSccMxKoCgdYdx9lLyXY" // Third secret and onwards
		}
	}

	// Function to simulate secret rotation
	rotateSecret := func() {
		atomic.AddInt32(&rotationCount, 1)
		RotateSecrets()
	}

	const exampleText = "The quick brown fox jumps over the lazy dog"
	algo := algorithm.SHA256
	originalSignature := Sign(algo, exampleText)

	// No Rotation: Signature should be valid
	t.Run("No Rotation", func(t *testing.T) {
		if !VerifySignature(algo, exampleText, originalSignature) {
			t.Errorf("Signature should be valid with zero rotations")
		}
	})

	// First Rotation: Signature should still be valid
	t.Run("First Rotation", func(t *testing.T) {
		rotateSecret()
		if !VerifySignature(algo, exampleText, originalSignature) {
			t.Errorf("Signature should be valid after one rotation")
		}
	})

	// Second Rotation: Signature should now be invalid
	t.Run("Second Rotation", func(t *testing.T) {
		rotateSecret()
		if VerifySignature(algo, exampleText, originalSignature) {
			t.Errorf("Signature should be invalid after two rotations")
		}
	})
}

func TestSignPanicOnEmptySecret(t *testing.T) {
	// Override randomString to return an empty string
	originalRandomString := internal.RandomStringReal
	internal.RandomString = func(length int) string {
		return ""
	}
	defer func() {
		// Restore the original randomString after the test
		internal.RandomString = originalRandomString

		// Check for panic
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty secret, but did not panic")
		}
	}()

	// Rotate secrets to set the current secret to an empty string
	RotateSecrets()

	// Call Sign with valid parameters, expecting a panic due to empty secret
	_ = Sign(algorithm.SHA256, "test text")
}
