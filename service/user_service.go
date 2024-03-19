package service

import (
	"fmt"

	"github.com/tcp-x/cd-core/sys/base"
	"github.com/tcp-x/cd-core/sys/user"
)

type CdRequest struct {
	Req string
}

type CdResponse struct {
	Resp string
}

type UserService struct{}

type User struct{}

/*
*
  - {
    "ctx": "Sys",
    "m": "User",
    "c": "User",
    "a": "Login",
    "dat": {
    "f_vals": [
    {
    "data": {
    "userName": "jondoo",
    "password": "iiii",
    "consumerGuid": "B0B3DA99-1859-A499-90F6-1E3F69575DCD"
    }
    }
    ],
    "token": ""
    },
    "args": null
    }
  - @param req
  - @param res
*/
func (t *User) Auth(req *CdRequest, resp *CdResponse) (string, error) {
	fmt.Println("UserService::Auth()/req:", req.Req)

	authResult, err := user.Auth(req.Req)
	user.
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	fmt.Println("authResult:", authResult)

	// get user and anon data
	// 1. convert req to struct
	reqStruct, err := base.JSONToICdRequest(req.Req)
	if err != nil {
		fmt.Println("Error:", err)
		return "", nil
	}

	// Accessing fields of MyStruct
	fmt.Println("Module:", reqStruct.M)
	fmt.Println("Dat:", reqStruct.Dat)

	// this.plData = this.b.getPlData(req);
	//     const q: IQuery = {
	//         // get requested user and 'anon' data/ anon data is used in case of failure
	//         where: [
	//             { userName: this.plData.userName },
	//             { userName: "anon" }
	//         ]
	//     };

	// connect to db and check validity of password
	// Auth input should have username and password

	// test if /tcp-x/user/session is accessible
	sid := user.SessID()
	fmt.Println("cd-user/Auth(): SessionID:", sid)

	// resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	resp.Resp = `{"state":"success", "data":[]}`
	return resp.Resp, nil
}
