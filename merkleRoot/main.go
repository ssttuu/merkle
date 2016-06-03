package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

func Reverse(s string) string {
	b := make([]byte, len(s));
	var j int = len(s) - 1;
	for i := 0; i <= j; i++ {
		b[j - i] = s[i]
	}

	return string(b);
}

func Hash2(a string, b string) string {
	hasher := sha256.Sum256([]byte(Reverse(b) + Reverse(a)))
	return hex.EncodeToString(hasher[:])
}

func MerkleRoot(hashes []string) string {
	if len(hashes) == 1 {
		return hashes[0]
	}

	newHashList := []string{}
	for i := 0; i < len(hashes); i += 2 {
		fmt.Println(i)
		newHashList = append(newHashList, Hash2(hashes[i], hashes[i + 1]))
	}

	if len(hashes) % 2 == 1 {
		lastItem := hashes[len(hashes) - 1]
		newHashList = append(newHashList, Hash2(lastItem, lastItem))
	}

	return MerkleRoot(newHashList)
}

func main() {
	txHashes := []string{"transaction1", "transaction2"}

	m := MerkleRoot(txHashes)
	fmt.Println(m)
}
