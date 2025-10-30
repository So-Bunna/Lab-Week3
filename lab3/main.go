package main

import (
	"flag"
	"fmt"
	"os"

	"lab3/utils/crack" // replace "yourmodule" with your module name
)

func main() {
	hashPtr := flag.String("hash", "", "Target SHA512 hash (hex).")
	wordlistPtr := flag.String("wordlist", "", "Path to wordlist file.")
	verbosePtr := flag.Bool("v", false, "Verbose: print each attempt.")
	flag.Parse()

	if *hashPtr == "" || *wordlistPtr == "" {
		fmt.Printf("Usage: %s -hash <sha512-hex> -wordlist <path> [-v]\n", os.Args[0])
		os.Exit(1)
	}

	found, err := crack.CrackSHA512(*hashPtr, *wordlistPtr, *verbosePtr, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if found == "" {
		fmt.Println("Password not found in the provided wordlist.")
	} else {
		fmt.Printf("Password cracked: %s\n", found)
	}
}
