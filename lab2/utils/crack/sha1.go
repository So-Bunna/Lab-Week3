package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hash returns the hex-encoded SHA-1 of the input string.
func SHA1Hash(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
