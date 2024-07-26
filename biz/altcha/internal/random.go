package internal

import (
	"crypto/rand"
	"math"
	"math/big"
	"strings"
)

var (
	RandomInt    = RandomIntReal
	RandomString = RandomStringReal
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits int64 = 6                    // 6 bits to represent a letter index
	letterIdxMask       = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax        = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var maxBigInt = big.NewInt(math.MaxInt64)

const alphabet = "0987654321ZYXWVUTSRQPONMLKJIHGFEDCBAzyxwvutsrqponmlkjihgfedcba"

func RandomStringReal_(length int) string {
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[RandomIntReal(0, len(alphabet))]
	}
	return string(b)
}

// RandomIntReal Int generates a random int between the given range
func RandomIntReal(minimum, maximum int) int {
	maxBigInt := big.NewInt(int64(maximum - minimum))
	number, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}
	return minimum + int(number.Int64())
}

// RandomStringReal String generates a random string of given length
func RandomStringReal(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := length-1, randomInt64(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randomInt64(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

func randomInt64() int64 {
	n, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
