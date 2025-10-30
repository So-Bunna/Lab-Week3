package main

import (
	"bufio"
	"flag"
	"fmt"
	"lab2/utils/crack"
	"log"
	"os"
	"strings"
)

func main() {
	wordlistPath := flag.String("wordlist", "", "path to wordlist file")
	target := flag.String("hash", "", "target sha1 hash to crack (hex lowercase)")
	verbose := flag.Bool("v", false, "verbose: print each attempt")
	limit := flag.Int("limit", 0, "optional: stop after this many attempts (0 = no limit)")
	flag.Parse()

	if *wordlistPath == "" || *target == "" {
		flag.Usage()
		os.Exit(2)
	}

	f, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatalf("failed to open wordlist: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	attempts := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		attempts++
		h := crack.SHA1Hash(line)

		if *verbose {
			// prints attempted password and its sha1
			fmt.Printf("[%d] trying: %q => %s\n", attempts, line, h)
		}

		if h == *target {
			fmt.Println("--------------------------------------------------")
			fmt.Printf("FOUND! password = %q\n", line)
			fmt.Printf("hash = %s\n", h)
			fmt.Printf("attempts = %d\n", attempts)
			fmt.Println("--------------------------------------------------")
			os.Exit(0)
		}

		if *limit > 0 && attempts >= *limit {
			fmt.Printf("Reached attempt limit (%d). Stopping.\n", *limit)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	fmt.Println("Not found in wordlist.")
}
