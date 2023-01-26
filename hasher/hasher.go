package hasher

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const (
	argonTime       = 3
	argonMemory     = 32768
	argonThreads    = 4
	ArgonKeyLength  = 32
	CondimentLength = 16
)

func GenerateCondiment() ([]byte, error) {
	condiment := make([]byte, CondimentLength)
	_, err := rand.Read(condiment)
	return condiment, err
}

func CreateHash(password []byte, pepper []byte) (salt []byte, hash []byte, err error) {
	salt, err = GenerateCondiment()
	if err != nil {
		return []byte{}, []byte{}, err
	}
	hash = argonWithDefaults(password, salt, pepper)
	return
}

func CheckPassword(password []byte, salt []byte, pepper []byte, correctHash []byte) bool {
	testHash := argonWithDefaults(password, salt, pepper)
	return bytes.Equal(testHash, correctHash)
}

func argonWithDefaults(password []byte, salt []byte, pepper []byte) []byte {
	return argon2.Key(
		argon2.Key(password, pepper, argonTime, argonMemory, argonThreads, ArgonKeyLength),
		salt,
		argonTime,
		argonMemory,
		argonThreads,
		ArgonKeyLength,
	)
}
