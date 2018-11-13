package main
// Exaample from: https://stevenwhite.com/building-a-rest-service-with-golang-2/

//Test:
//    curl http://localhost:3000/user/1
// should produce:  {"name":"Bob Smith","gender":"male","age":50,"id":"1"}

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/swhite24/go-rest-tutorial/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController()

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

