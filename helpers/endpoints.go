package helpers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)


type EncryptService interface{
	Encrypt(context.Context, string, string) (string, error)
	Decrypt(context.Context, string, string) (string, error)
}

func MakeEncryptEndpoint(srv EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EncryptRequest)
		message, err := srv.Encrypt(ctx, req.Key, req.Text)

		if err != nil {
			return EncryptResponse{message, err.Error()}, nil
		}

		return EncryptResponse{message, ""}, nil
	}
}
func MakeDecryptEndpoint(srv EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DecryptRequest)
		message, err := srv.Decrypt(ctx, req.Key, req.Message)

		if err != nil {
			return DecryptResponse{message, err.Error()}, nil
		}

		return DecryptResponse{message, ""}, nil
	}
}