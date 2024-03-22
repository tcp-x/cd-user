/*
RPC Server with following procedures:

    Retrieve a To-Do
    Make a To-Do
    Edit a To-Do
    Delete a To-Do

Ref:
- https://medium.com/@OmisNomis/creating-an-rpc-server-in-go-3a94797ab833
- https://gist.github.com/OmisNomis/7dd87d72179e98c560ab735b4dc2039f
- https://github.com/OmisNomis/CdResponseRPC

---------------------
Go RPC Requirements:
---------------------

The net/rpc package stipulates that only methods that satisfy the following criteria will be made available for remote access; other methods will be ignored.

    The method’s type is exported.
    The method is exported
    The method has two arguments, both exported (or builtin types).
    The method’s second argument is a pointer
    The method has return type error

// Example:
// From the appliction todo-withou-rpc, we have:

func MakeCdResponse(todo CdResponse) CdResponse {
	todoSlice = append(todoSlice, todo)
	return todo
}

// First, we need to make a method type and make sure it’s exported (make sure the name is capitalised)

type Tasks int

// Then we use it as a method receiver.

func (t *Tasks) MakeCdResponse(todo CdResponse) CdResponse { }

// The method is already exported (because the name, MakeCdResponse, is capitalised) so the next thing to do is
// change the methods arguments. From the RPC documentation:
// 		"The method’s first argument represents the arguments provided by the caller; the second argument
// 		represents the result parameters to be returned to the caller. The method’s return value, if non-nil,
// 		is passed back as a string that the client sees as if created by errors.New. If an error is returned,
// 		the reply parameter will not be sent back to the client."

func (t *Tasks) MakeCdResponse(todo CdResponse, reply *CdResponse) CdResponse { }

// Now the method has two arguments, todo and reply, and both have builtin or exported types. Notice that the reply type is CdResponse,
// which is what we expect to return to the client if not an error.
// // The first argument (todo) will be provided by the caller; the second argument (reply) represents what will be returned to the caller.
// We have also covered the next requirement, that the method’s second argument is a pointer.
// That leaves one thing left to do; change the return type to be error.

func (t *Tasks) MakeCdResponse(todo CdResponse, reply *CdResponse) error { }


// Our method now follows the RPC rules but, if you were to try and run it the same way as before, you will get errors.
// This is due to two things we need to refactor:

//     The method body. We are currently returning todo, but the method is expecting an error to be returned.
//     The method execution. Firstly we can no longer call it like a function (MakeCdResponse(finishApp)), because it’s a method,
// 	and secondly we need to amend the arguments, because it now expects two.

// Let’s address the method body first. A quick reminder about the Go requirements for an RPC Method:

//     The method’s first argument represents the arguments provided by the caller; the second argument represents the result
// 	parameters to be returned to the caller. The method’s return value, if non-nil, is passed back as a string that the client sees as if created by errors.New. If an error is returned, the reply parameter will not be sent back to the client.

// The important parts, that outline what we need to change in our method body, are:

//     The second argument represents the result parameters to be returned to the caller.
//     The methods return value, if non-nil, is passed back as a string that the client sees as if created by errors.New
//     If an error is returned, the reply parameter will not be sent back to the client.

// This means that, instead of returning a value, we use a pointer to the reply parameter and return nil to indicate no errors.
// If we do want to return an error we just return it and the reply parameter will not be sent back to the client.

func (t *Task) MakeCdResponse(todo CdResponse, reply *CdResponse) error {
	todoSlice = append(todoSlice, todo)
	*reply = todo
	return nil
}

// Now lets fix how we’re calling the method, and the parameters we’re passing. First, in our main function, we need to
// register a new object that we can call our method on.

func main() {
	task := new(Task)
}

// We now have an object (task) that we can use to call our methods. We can now change our method invocation from this:

MakeCdResponse(finishApp)

// To this:

task.MakeCdResponse(finishApp)

// We now only have one more problem to fix, and that’s our parameters.

func main() {
	task := new(Task)
	var err error
	finishApp := CdResponse{"Finish App", "Started"}
	var makeReply CdResponse
	err = task.MakeCdResponse(finishApp, &makeReply)
	if err != nil {
		log.Fatal("Issue making CdResponse: ", err)
	}
}

// We have now declared three new variables:

//     err that is of type error
//     finishApp that is of type CdResponse
//     makeReply that is of type CdResponse

// We then call our MakeCdResponse method with the finishApp parameter like before, and an additional parameter &makeReply.
// The & denotes that this is a pointer to our makeReply variable. This means that the makeReply variable will be updated
// with the result of our method. Remember the Go RPC spec states the method’s first argument represents the arguments
// provided by the caller; the second argument represents the result parameters to be returned to the caller which
// is what we have defined.

// Because the MakeCdResponse method has a return type of error, we can use err = MakeCdResponse(…)
// to determine whether an error was returned or not (in our case it’s nil).

*/

package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/tcp-x/cd-core/sys/base"
	"github.com/tcp-x/cd-core/sys/user"
)

// Make a new CdResponse type that is a typed collection of fields
// (Title and Status), both of which are of type string
type CdResponse struct {
	AppState CdAppState
	Data     RespData
}

type CdAppState struct {
	Success bool
	Info    string
	Sess    string
	Cache   string
	SConfig string
}

type RespData struct {
	Data           []User
	RowsAffected   int
	NumberOfResult int
}

type CdRequest struct {
	Ctx string
	M   string
	C   string
	A   string
	Dat FValDat
}

type FValDat struct {
	F_vals FValItem
	Token  string
}

type FValItem struct {
	Data User
}

type User struct {
	UserName string
	Password string
	Email    string
}

// Declare variable 'user' that is a slice made up of
// type CdResponse items
var userStore []User

type UserController int

var serverPort = 12345

var logger base.Logger

// var logIndex = 0

// As per documentation:
//
//	The method’s first argument represents the arguments provided by the caller; the second
//	argument represents the result parameters to be returned to the caller. The method’s return
//	value, if non-nil, is passed back as a string that the client sees as if created by errors.New.
//	If an error is returned, the reply parameter will not be sent back to the client.
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
func (t *UserController) Auth(req user.CdRequest, resp *user.CdResponse) error {
	logger.LogInfo("Starting rpc server UserController::Auth()...")
	logger.LogInfo("Starting rpc server UserController::Auth()/req:" + fmt.Sprint(req))
	var err error
	authResult := user.Auth(req)
	logger.LogInfo("UserController::Auth()/authResult" + fmt.Sprint(authResult))
	*resp = authResult
	return err
}

// GetUser takes a string type and returns a ToDo
func (t *UserController) GetUser(req CdRequest, resp *CdResponse) error {
	var found []User
	// Range statement that iterates over todoArray
	// 'v' is the value of the current iterateee
	for _, u := range userStore {
		if u.Email == req.Dat.F_vals.Data.Email {
			found = append(found, u)
		}
	}
	var appState = CdAppState{true, "", "", "", ""}
	var appData = RespData{Data: found, RowsAffected: 0, NumberOfResult: 1}
	*resp = CdResponse{AppState: appState, Data: appData}
	return nil
}

func (t *UserController) Create(req CdRequest, resp *CdResponse) error {
	// save user to user strore
	userStore = append(userStore, req.Dat.F_vals.Data)
	var appState = CdAppState{true, "", "", "", ""}
	var appData = RespData{Data: nil, RowsAffected: 1, NumberOfResult: 0}
	*resp = CdResponse{AppState: appState, Data: appData}
	return nil
}

// Search by email and edit by email
func (t *UserController) EditPassword(req CdRequest, resp *CdResponse) error {
	var edited User
	// 'i' is the index in the array and 'v' the value
	for i, u := range userStore {
		if u.Email == req.Dat.F_vals.Data.Email {
			edited = User{userStore[i].UserName, userStore[i].Email, req.Dat.F_vals.Data.Password}
			userStore[i] = edited
		}
	}
	// edited will be the edited ToDo or a zeroed ToDo
	var appState = CdAppState{true, "Edit was successfull", "", "", ""}
	var appData = RespData{Data: nil, RowsAffected: 1, NumberOfResult: 0}
	*resp = CdResponse{AppState: appState, Data: appData}
	return nil
}

// Search by email and delete user
func (t *UserController) Delete(req CdRequest, resp *CdResponse) error {
	// 'i' is the index in the array and 'v' the value
	for i, u := range userStore {
		if u.Email == req.Dat.F_vals.Data.Email {
			userStore = append(userStore[:i], userStore[i+1:]...)
		}
	}
	var appState = CdAppState{true, "Delete was successfull", "", "", ""}
	var appData = RespData{Data: nil, RowsAffected: 1, NumberOfResult: 0}
	*resp = CdResponse{AppState: appState, Data: appData}
	return nil
}

// func setCdRequest(newUser User) CdRequest {
// 	// newUser := User{"karl", "secret", "karl@emp.net"}
// 	fvalItem := FValItem{newUser}
// 	fvalDat := FValDat{fvalItem, ""}
// 	return CdRequest{"Sys", "UserModule", "UserController", "Create", fvalDat}
// }

func main() {
	logger.LogInfo("starting cd-server-v2...")
	userHandle := new(UserController)
	// var err error

	// // -------------------------------
	// // INSERT 3 USERS
	// // -------------------------------
	// req1 := setCdRequest(User{"karl", "secret", "karl@emp.net"})
	// req2 := setCdRequest(User{"zog", "zigdin", "zog@emp.net"})
	// req3 := setCdRequest(User{"jdoe", "admin", "jdoe@emp.net"})
	// var cdResp CdResponse
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// err = userHandle.Create(req1, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue registering new User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// err = userHandle.Create(req2, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue registering new User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// err = userHandle.Create(req3, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue registering new User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)

	// // -------------------------------
	// // GET USER
	// // -------------------------------
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req4 := setCdRequest(User{"", "", "jdoe@emp.net"})
	// err = userHandle.GetUser(req4, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue getting User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	// // -------------------------------
	// // DELETE USER
	// // -------------------------------
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req5 := setCdRequest(User{"", "", "jdoe@emp.net"})
	// err = userHandle.Delete(req5, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue deleting User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	// // -------------------------------
	// // EDIT USER
	// // -------------------------------
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)
	// req6 := setCdRequest(User{"", "xxyy", "zog@emp.net"})
	// err = userHandle.EditPassword(req6, &cdResp)
	// if err != nil {
	// 	log.Fatal("Issue editing User: ", err)
	// }
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+".", cdResp)
	// logIndex++
	// log.Println(strconv.Itoa(logIndex)+". userStore: ", userStore)

	///////////////////////////////////////////////////////////////////////////////////////////////

	// -------------------------------
	// START RPC SERVER
	// -------------------------------

	// task := new(Task)
	// Publish the receivers methods
	err := rpc.Register(userHandle)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":"+strconv.Itoa(serverPort))
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", serverPort)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
