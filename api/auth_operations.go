package api

import (
	"context"
	"authservice/pkg/protobuff/api/auth"
	"authservice/pkg/protobuff/api/user"
	"authservice/pkg/protobuff/service"
	"authservice/pkg/models"
)

var allowedFunc = map[string]bool{
	"/authservice.service.Auth/SignUp":   true,
}

// AuthFuncOverride This will bypass on method matching allowedFunc
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if allowedFunc[fullMethodName] {
		return ctx, nil
	}
	return s.Auth.VerifyCredentials(ctx)
}

// Ping RPC for checking connection to service.
func (s *Server) Ping(ctx context.Context, in *service.Empty) (*service.Empty, error) {
	return &service.Empty{}, nil
}

func (s *Server) SignUp(ctx context.Context,  req *user.CreateUserRequest) (*auth.AuthResponse, error) {
	userModel := &models.User {
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userModel, err := userModel.CreateUser(s.DB)
	if err != nil {
		return nil, err
	}

	token, err := s.Auth.EncodeToken(userModel)
	if err != nil {
		// ctx.Logger.Error().Err(err).Msg("unable to encode jwt")
		return nil, err
	}

	return &auth.AuthResponse{
		Id:       int32(userModel.ID),
		Token:    token.AccessToken,
		Email:	  userModel.Email,
	}, nil
}

func (s *Server) Login(context.Context, *auth.AuthRequest) (*auth.AuthResponse, error) {
	return nil, nil
}