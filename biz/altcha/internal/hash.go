package internal

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/tundrawork/powcha/biz/altcha/algorithm"
	"hash"
	"strconv"
	"sync"
)

var hasherPoolSha256 = sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

var hasherPoolSha384 = sync.Pool{
	New: func() interface{} {
		return sha512.New384()
	},
}

var hasherPoolSha512 = sync.Pool{
	New: func() interface{} {
		return sha512.New()
	},
}

func GetHasher(algo algorithm.Algorithm) (hasher hash.Hash, put func()) {
	switch algo {
	case algorithm.SHA256:
		hasher = hasherPoolSha256.Get().(hash.Hash)
		put = func() {
			hasher.Reset()
			hasherPoolSha256.Put(hasher)
		}
	case algorithm.SHA384:
		hasher = hasherPoolSha384.Get().(hash.Hash)
		put = func() {
			hasher.Reset()
			hasherPoolSha384.Put(hasher)
		}
	case algorithm.SHA512:
		hasher = hasherPoolSha512.Get().(hash.Hash)
		put = func() {
			hasher.Reset()
			hasherPoolSha512.Put(hasher)
		}
	default:
		panic("unknown hashing algorithm")
	}
	return hasher, put
}

func GenerateHash(algo algorithm.Algorithm, salt string, number int) string {
	hasher, put := GetHasher(algo)
	defer put()
	hasher.Write([]byte(salt))
	hasher.Write([]byte(strconv.Itoa(number)))
	return hex.EncodeToString(hasher.Sum(nil))
}
