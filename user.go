package main

import "fmt"

func Login(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := data
	return resp
}

/*
- Consider auth using public-key signed by corpdesk CA
*/
func Auth(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
- Present varifiable email address
*/
func Register(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}
