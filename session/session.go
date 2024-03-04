package main

import (
	"fmt"
)

func SessionInit(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := data
	return resp
}
