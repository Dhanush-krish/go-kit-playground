package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// service interface
// service.go
type StringService interface {
	Length(string) int
	Upper(string) string
}

// service implementation
// logic.go
type stringService struct{}

// business logic
func (stringService) Length(str string) int {
	return len(str)
}

func (stringService) Upper(str string) string {
	fmt.Println("I'm in Upper")
	return strings.ToUpper(str)
}

// rpc functions
// reqresp.go
type lengthRequest struct {
	S string `json:"s"`
}

type upperRequest struct {
	S string `json:"s"`
}

type lengthResponse struct {
	V int `json:"v"`
}

type upperResponse struct {
	S string `json:"s"`
}

// endpoints
// endpoints.go
func makeLengthEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("I'm in makeLengthEndpoint")
		req := request.(lengthRequest)
		v := svc.Length(req.S)
		return lengthResponse{v}, nil
	}
}

func makeUpperEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(upperRequest)
		fmt.Println("I'm in makeUpperEndpoint")
		converted := svc.Upper(req.S)
		return upperResponse{converted}, nil
	}
}

func main() {

	fmt.Println("Hello go kit")
	svc := stringService{}

	lengthHandler := httptransport.NewServer(
		makeLengthEndpoint(svc),
		decodeLengthRequest,
		encodeLengthRequest,
	)

	upperHandler := httptransport.NewServer(
		makeUpperEndpoint(svc),
		decodeUpperRequest,
		encodeUpperRequest)

	http.Handle("/count", lengthHandler)
	http.Handle("/upper", upperHandler)

	log.Fatal(http.ListenAndServe(":9090", nil))

}

func decodeLengthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request lengthRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil

}

func encodeLengthRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUpperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req_payload upperRequest
	if err := json.NewDecoder(r.Body).Decode(&req_payload); err != nil {
		return nil, err
	}
	fmt.Println("I'm in decodeUpperRequest")

	return req_payload, nil
}

func encodeUpperRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("I'm in encodeUpperRequest")
	return json.NewEncoder(w).Encode(response)
}
