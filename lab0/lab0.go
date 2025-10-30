package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

// Function to compute hashes
func computeHashes(input string) map[string]string {
	hashes := make(map[string]string)

	// MD5
	md5Hash := md5.Sum([]byte(input))
	hashes["MD5"] = hex.EncodeToString(md5Hash[:])

	// SHA1
	sha1Hash := sha1.Sum([]byte(input))
	hashes["SHA1"] = hex.EncodeToString(sha1Hash[:])

	// SHA256
	sha256Hash := sha256.Sum256([]byte(input))
	hashes["SHA256"] = hex.EncodeToString(sha256Hash[:])

	// SHA512
	sha512Hash := sha512.Sum512([]byte(input))
	hashes["SHA512"] = hex.EncodeToString(sha512Hash[:])

	// SHA3-256
	sha3Hash := sha3.Sum256([]byte(input))
	hashes["SHA3"] = hex.EncodeToString(sha3Hash[:])

	return hashes
}

func proofMe(txt1, txt2 string) {
	fmt.Println("======== Name + Hashing Program ========")

	hashes1 := computeHashes(txt1)
	hashes2 := computeHashes(txt2)

	for algo, hash1 := range hashes1 {
		hash2 := hashes2[algo]
		match := "No Match!"
		if strings.EqualFold(hash1, hash2) {
			match = "Match!"
		}
		fmt.Printf("Hash (%s):\nOutput A: %s\nOutput B: %s\n=> %s\n\n", algo, hash1, hash2, match)
	}
}

func main() {
	var input1, input2 string

	fmt.Print("Please input value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Please input value 2: ")
	fmt.Scanln(&input2)

	proofMe(input1, input2)
}
