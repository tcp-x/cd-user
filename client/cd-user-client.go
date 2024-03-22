package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"

	"github.com/tcp-x/cd-core/sys/base"
	"github.com/tcp-x/cd-core/sys/user"
)

// // Make a new CdResponse type that is a typed collection of fields
// // (Title and Status), both of which are of type string
// type CdResponse struct {
// 	AppState CdAppState
// 	Data     RespData
// }

// type CdAppState struct {
// 	Success bool
// 	Info    string
// 	Sess    string
// 	Cache   string
// 	SConfig string
// }

// type RespData struct {
// 	Data           []User
// 	RowsAffected   int
// 	NumberOfResult int
// }

// type CdRequest struct {
// 	Ctx string
// 	M   string
// 	C   string
// 	A   string
// 	Dat FValDat
// }

// type FValDat struct {
// 	F_vals FValItem
// 	Token  string
// }

// type FValItem struct {
// 	Data User
// }

// type User struct {
// 	UserName string
// 	Password string
// 	Email    string
// }

// Declare variable 'user' that is a slice made up of
// type CdResponse items
var userStore []user.User

type UserController int

var serverPort = 12345
var logger base.Logger
var logIndex = 0

type EditToDo struct {
	Title, NewTitle, NewStatus string
}

func setCdRequest(newUser user.User) user.CdRequest {
	// newUser := User{"karl", "secret", "karl@emp.net"}
	fvalItem := user.FValItem{newUser}
	fvalDat := user.FValDat{fvalItem, ""}
	return user.CdRequest{"Sys", "UserModule", "UserController", "Create", fvalDat}
}

func main() {
	// userHandle := new(UserController)
	var err error
	var cdResp user.CdResponse

	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:"+strconv.Itoa(serverPort))
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// -------------------------------
	// INSERT 3 USERS
	// -------------------------------
	req1 := setCdRequest(user.User{0, "", "karl", "secret", "karl@emp.net", 0, "", false, time.Now(), "", "", "", "", 0, 0, false, "", "", 0, 0})
	req2 := setCdRequest(user.User{0, "", "zog", "zigdin", "zog@emp.net", 0, "", false, time.Now(), "", "", "", "", 0, 0, false, "", "", 0, 0})
	req3 := setCdRequest(user.User{0, "", "jdoe", "admin", "jdoe@emp.net", 0, "", false, time.Now(), "", "", "", "", 0, 0, false, "", "", 0, 0})
	// client.Call("UserController.Create", req1, &cdResp)
	// client.Call("UserController.Create", req2, &cdResp)
	// client.Call("UserController.Create", req3, &cdResp)
	// var cdResp CdResponse
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// err = userHandle.Create(req1, &cdResp)
	err = client.Call("UserController.Create", req1, &cdResp)
	if err != nil {
		log.Fatal("Issue registering new User: ", err)
	}
	logIndex++
	log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	err = client.Call("UserController.Create", req2, &cdResp)
	if err != nil {
		log.Fatal("Issue registering new User: ", err)
	}
	logIndex++
	log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// err = userHandle.Create(req3, &cdResp)
	err = client.Call("UserController.Create", req3, &cdResp)
	if err != nil {
		log.Fatal("Issue registering new User: ", err)
	}
	logIndex++
	log.Println(strconv.Itoa(logIndex)+".", cdResp)

	// // -------------------------------
	// // GET USER
	// // -------------------------------
	// // logIndex++
	// // log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req4 := setCdRequest(user.User{"", "", "jdoe@emp.net"})
	// // err = userHandle.GetUser(req4, &cdResp)
	// err = client.Call("UserController.GetUser", req4, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue getting User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// // logIndex++
	// // log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	// // -------------------------------
	// // DELETE USER
	// // -------------------------------
	// // logIndex++
	// // log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req5 := setCdRequest(User{"", "", "jdoe@emp.net"})
	// // err = userHandle.Delete(req5, &cdResp)
	// err = client.Call("UserController.GetUser", req5, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue deleting User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// // logIndex++
	// // log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	// // -------------------------------
	// // EDIT USER
	// // -------------------------------
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req6 := setCdRequest(User{"", "xxyy", "zog@emp.net"})
	// // err = userHandle.EditPassword(req6, &cdResp)
	// err = client.Call("UserController.EditPassword", req6, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue editing User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)

	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	// -------------------------------
	// AUTH USER
	// -------------------------------
	req7 := setCdRequest(user.User{0, "", "karl", "secret", "", 0, "", false, time.Now(), "", "", "", "", 0, 0, false, "", "", 0, 0})
	// err = userHandle.EditPassword(req6, &cdResp)
	err = client.Call("UserController.Auth", req7, &cdResp)
	if err != nil {
		log.Fatal("Issue authenticating User: ", err)
	}
	logIndex++
	log.Println(strconv.Itoa(logIndex)+".", cdResp)

	respJStr, err := json.Marshal(cdResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(respJStr))
	// var r string = string(respJStr)
	logIndex++
	logger.LogInfo(strconv.Itoa(logIndex) + ". Response: " + string(respJStr))
}
