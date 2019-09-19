package auth

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"authservice/pkg/models"
)

// Token response body for JWT
type Token struct {
	AccessToken string `json:"access_token"`
}

// IAuth is the interface that wraps the auth method.
type IAuth interface {
	// EncodeToken creates a token for a user argument for subsequent call verification.
	EncodeToken(user *models.User) (*Token, error)
	// VerifyCredentials called on each and every request made.
	// It decodes the token and puts userID and email as principal in the context.
	VerifyCredentials(ctx context.Context) (context.Context, error)
}

// UnaryServerInterceptor registers as an unary server interceptor
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(nil)
}

func (t Token) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.AccessToken,
	}, nil
}

func (Token) RequireTransportSecurity() bool {
	return true
}