package algorithm

type Algorithm int

const (
	UnknownAlgorithm Algorithm = iota
	SHA256
	SHA384
	SHA512
)

func (algorithm Algorithm) String() string {
	switch algorithm {
	case SHA256:
		return "SHA-256"
	case SHA384:
		return "SHA-384"
	case SHA512:
		return "SHA-512"
	default:
		return "Unknown"
	}
}

func AlgorithmFromString(algo string) (Algorithm, bool) {
	switch algo {
	case "SHA-256":
		return SHA256, true
	case "SHA-384":
		return SHA384, true
	case "SHA-512":
		return SHA512, true
	default:
		return UnknownAlgorithm, false
	}
}
