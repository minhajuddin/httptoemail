package main

import (
	hasher "crypto/sha512"
	"encoding/hex"
	"io"
)

func hash(input string) string {
	h := hasher.New()
	io.WriteString(h, input)
	return hex.EncodeToString(h.Sum(nil))
}
