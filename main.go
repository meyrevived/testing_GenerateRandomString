package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		TestGenRanStrLoopCompare()
		TestGenRanStrSortedSliceIndex()
	}
}

func loopCompare(s []string, target string) bool {
	// turns out simple comparison operators function the fastest in Go string comparison...
	for _, str := range s {
		if str == target {
			return true
		}
	}
	return false
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func TestGenRanStrLoopCompare() {
	var testRounds = 10000

	var ideticalStrings = 0
	var errorCount = 0
	startTime := time.Now()
	randomStrings := make([]string, 0, testRounds)
	for i := 0; i < testRounds; i++ {
		randStr, err := GenerateRandomString(20)
		if err != nil {
			errorCount++
		} else if loopCompare(randomStrings, randStr) {
			ideticalStrings++
		} else {
			randomStrings = append(randomStrings, randStr)

		}
	}
	runTime := time.Since(startTime)
	fmt.Printf("LoopCompare tested GenerateRandomString %d random strings in %s\n", testRounds, runTime)
}

func TestGenRanStrSortedSliceIndex() {
	var testRounds = 10000

	var ideticalStrings = 0
	var errorCount = 0
	startTime := time.Now()
	randomStrings := make([]string, 0, testRounds)
	for i := 0; i < testRounds; i++ {
		randStr, err := GenerateRandomString(20)
		if err != nil {
			errorCount++
		} else if slices.Index(randomStrings, randStr) != -1 {
			ideticalStrings++
		} else {
			randomStrings = append(randomStrings, randStr)
			sort.Strings(randomStrings)
		}
	}
	runTime := time.Since(startTime)
	fmt.Printf("SortedSliceIndex tested GenerateRandomString %d random strings in %s\n", testRounds, runTime)
}
