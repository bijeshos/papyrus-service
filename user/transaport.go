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
		v, err := svc.AddUser(req.S)
		if err != nil {
			return AddUserResponse{v, err.Error()}, nil
		}
		return AddUserResponse{v, ""}, nil
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
	S string `json:"s"`
}

type AddUserResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
