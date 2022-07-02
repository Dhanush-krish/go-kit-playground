package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	http_transport "github.com/go-kit/kit/transport/http"
)

type UserInterface interface {
	greetUser(string) string
}

//class
type User struct{}

//method
func (u User) greetUser(name string) string {

	return "Welcome " + name + "!"
}

type greetResponse struct {
	Message string `json:"message"`
}

type greetRequest struct {
	Name string `json:"name"`
}

func makeGreetUserEndpoint(u UserInterface) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(greetRequest)
		resp := u.greetUser(req.Name)

		return greetResponse{resp}, nil

	}

}

func main() {

	fmt.Println("*********Greeter API Using Go-kit*********")
	fmt.Println("")

	//create service
	svc := User{}
	greetUser := http_transport.NewServer(
		makeGreetUserEndpoint(svc),
		decodeGreetUserRequest,
		encodeGreetUserResponse)

	//http handler
	http.Handle("/greet", greetUser)

	//http listener
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func decodeGreetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req_body greetRequest

	if err := json.NewDecoder(r.Body).Decode(&req_body); err != nil {
		return nil, err
	}

	return req_body, nil
}

func encodeGreetUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
