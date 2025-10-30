package crack

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// CrackSHA512 tries to find the plaintext password for the given target SHA512 hex hash
// wordlistPath: path to a wordlist file (one candidate per line).
// targetHex: the target SHA512 hex string (case-insensitive).
// verbose: if true, prints attempts to w (usually os.Stdout).
// Returns the found password (or empty string if not found) and error.
func CrackSHA512(targetHex string, wordlistPath string, verbose bool, w io.Writer) (string, error) {
	target := strings.ToLower(strings.TrimSpace(targetHex))
	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lineNum int
	for scanner.Scan() {
		lineNum++
		pw := strings.TrimSpace(scanner.Text())
		if pw == "" {
			continue
		}

		sum := sha512.Sum512([]byte(pw))
		digest := hex.EncodeToString(sum[:])

		if verbose {
			// Print attempt: password -> digest
			fmt.Fprintf(w, "%d: %s -> %s\n", lineNum, pw, digest)
		}

		if digest == target {
			// Found
			fmt.Fprintf(w, "FOUND: password = %s (line %d)\n", pw, lineNum)
			return pw, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	return "", nil
}
