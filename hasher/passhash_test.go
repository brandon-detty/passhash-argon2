package hasher

import (
	"bytes"
	"testing"
)

func TestCreateAndVerifyPassword(t *testing.T) {
	password := []byte("password123")
	pepper, err := GenerateCondiment()
	if err != nil {
		t.Error("failed to generate salt", err)
	}
	salt, hash, err := CreateHash(password, pepper)
	if err != nil {
		t.Error("failed creating hash", err)
	}

	if !CheckPassword(password, salt, pepper, hash) {
		t.Error("CheckPassword() failed with result of CreateHash()")
	}
}

func TestSaltAndHashUniqueness(t *testing.T) {
	password := []byte("password123")
	pepper, err := GenerateCondiment()
	if err != nil {
		t.Error("failed to generate salt", err)
	}

	salt, hash, err := CreateHash(password, pepper)
	if err != nil {
		t.Error("failed creating hash", err)
	}
	salt2, hash2, err := CreateHash(password, pepper)
	if err != nil {
		t.Error("failed creating hash", err)
	}

	if bytes.Equal(salt, salt2) {
		t.Error("salt was reused by CreateHash()")
	}
	if bytes.Equal(hash, hash2) {
		t.Error("hash was reused by CreateHash()")
	}
}
