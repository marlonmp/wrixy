package service

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

const sidLen = 32

// this function generates a random session id
func RandomSID() string {
	buf := make([]byte, sidLen)

	io.ReadFull(rand.Reader, buf)

	return base64.RawURLEncoding.EncodeToString(buf)
}
