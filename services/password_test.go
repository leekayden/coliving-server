package services

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	bytes1, err := GenerateByteString()
	if err != nil {
		fmt.Println(err)
	}
	bytes2, err := GenerateByteString()
	if err != nil {
		fmt.Println(err)
	}
	if bytes.Equal(bytes1, bytes2) {
		t.Fatalf("Both strings are the same, should be different.")
	}
}

func TestHashPassword(t *testing.T) {
	password := "Password123!"
	hashed1, err := HashAndSaltPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	hashed2, err := HashAndSaltPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	if hashed1 == hashed2 {
		t.Fatalf("Both strings are the same, should be different.")
	}
}
