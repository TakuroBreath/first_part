package main

import "fmt"

type SecretString struct {
	secret string
}

func NewSecretString(s string) SecretString {
	return SecretString{
		secret: s,
	}
}

func (s SecretString) Value() string {
	return s.secret
}

func (s SecretString) String() string {
	return "******"
}

func main() {
	value := "sensetive data"
	
	s := NewSecretString(value)
	
	fmt.Print(s)
}
