package test

import (
	"fair-ticket-be-tutorial/cryptoo"
	"fmt"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	v := cryptoo.VerifySignature("create project",
		"0xce6d2fd413ff2081095cbbe8d8764521e4f798c69011856dfd4d0089eb0193ab4e4efc97fb1311b286f7f4f15f44907724c5ce7363e6d1ebc0d87b122834dbbe1c",
		"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println(v)
}
