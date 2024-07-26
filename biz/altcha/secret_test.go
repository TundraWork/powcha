package altcha

import (
	"github.com/tundrawork/powcha/biz/altcha/internal"
	"sync"
	"sync/atomic"
	"testing"
)

func TestGetSecrets(t *testing.T) {
	// Both secrets should never be empty
	current, previous := GetSecrets()
	t.Logf("Current secret: %s", current)
	t.Logf("Previous secret: %s", previous)
	if current == "" || previous == "" {
		t.Errorf("One or more secrets were empty")
	}
}

func TestRotateSecrets(t *testing.T) {
	// Setup: Override randomString and track calls to generate different secrets
	var callCount int32
	internal.RandomString = func(length int) string {
		switch atomic.AddInt32(&callCount, 1) {
		case 1:
			return "0V5xzYiSFmY1swbbkwIoAgbWaiw7yJvZ" // First secret
		case 2:
			return "1K7xwZjTHfM2tRbbLwJnBgbXcw8zKwXW" // Second secret
		default:
			return "2L8ywAkUIgN3uSccMxKoCgdYdx9lLyXY" // Third secret and onwards
		}
	}
	RotateSecrets() // Initial rotation to set up the above sequence

	originalCurrent, _ := GetSecrets()
	RotateSecrets()
	newCurrent, newPrevious := GetSecrets()

	if newCurrent == originalCurrent || newPrevious != originalCurrent {
		t.Logf("Original current secret: %s", originalCurrent)
		t.Logf("New current secret: %s", newCurrent)
		t.Errorf("Secrets were not rotated correctly")
	}
}

func TestConcurrencySafety(t *testing.T) {
	t.Logf("Testing concurrency safety. You should run this using the race detector.")

	internal.RandomInt = internal.RandomIntReal       // Reset randomInt to use the real function
	internal.RandomString = internal.RandomStringReal // Reset randomString to use the real function

	var wg sync.WaitGroup
	iterations := 100 // Number of concurrent calls to RotateSecrets

	// Use a WaitGroup to wait for all goroutines to finish
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			RotateSecrets() // Concurrently call RotateSecrets
		}()
	}

	wg.Wait() // Wait for all goroutines to complete

	// After all rotations, check if the secrets are valid
	current, previous := GetSecrets()
	if current == "" || previous == "" {
		t.Errorf("Secrets are empty after concurrent rotation")
	}
	if current == previous {
		t.Errorf("Current and previous secrets should not be the same")
	}
}
