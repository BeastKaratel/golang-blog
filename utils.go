package main

import (
	"fmt"
	"math/rand"
)

//GenerateID - random generate id for posts
func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
