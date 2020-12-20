package user

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeAdduserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		userId, err := svc.AddUser(req.FirstName, req.LastName)
		if err != nil {
			return AddUserResponse{userId, err.Error()}, nil
		}
		return AddUserResponse{userId, ""}, nil
	}
}

func DecodeAdduserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeAddUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response AddUserResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type AddUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AddUserResponse struct {
	UserId int    `json:"userId"`
	Err    string `json:"err,omitempty"`
}
