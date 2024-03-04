package main

import (
	"fmt"

	"github.com/tcp-x/cd-core/sys/base"
	"github.com/tcp-x/cd-core/sys/user"
)

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
	modReq := `{
		           "ctx": "Sys",
		           "m": "Moduleman",
		           "c": "CdObj",
		           "a": "GetCdObj",
		           "dat": {
		               "f_vals": [
		                   {
		                       "query": {
		                           "where": {"cdObjId": 45763}
		                       }
		                   }
		               ],
		               "token": "08f45393-c10e-4edd-af2c-bae1746247a1"
		           },
		           "args": null
		       }`

	// test access to moduleman plugin
	// access is via base.ExecPlug() which should be method available at the base
	// Example: var resp string = base.ExecPlug(args[0])
	fmt.Println("***Starting ExecPlug() at user:")
	modResp, err := base.ExecPlug(modReq)
	if err != nil {
		fmt.Println("Error executing user plugin:", err)
	}
	fmt.Println("Moduleman response:", modResp)

	// connect to db and check validity of password
	// Auth input should have username and password

	// test if /tcp-x/user/session is accessible
	sid := user.SessID()
	fmt.Println("cd-user/Auth(): SessionID:", sid)

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
