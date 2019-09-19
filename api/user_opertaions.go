package api

import (
	"context"
	"authservice/pkg/protobuff/api/user"
	"authservice/pkg/protobuff/service"
	"authservice/pkg/models"
	"authservice/pkg/utils"
	"strconv"
)


func (s *Server) GetUserDetails(ctx context.Context, req *service.Empty) (*user.UserDetails, error) {
	mdKeyUserID, err := utils.MDGetValue(ctx, utils.MDKeyUserID)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{}
	userID, err := strconv.ParseUint(mdKeyUserID, 10, 32)
	if err != nil {
		return nil, err
	}
	userModel, err = userModel.FindUserByID(s.DB, uint32(userID))
	if err != nil {
		return nil, err
	}

	return &user.UserDetails{
		Id:       int32(userModel.ID),
		Name:     userModel.Name,
		Email:	  userModel.Email,
	}, nil
}

func (s *Server) UpdateUserDetails(ctx context.Context, req *user.UpdateUserRequest) (*user.UserDetails, error) {
	mdKeyUserID, err := utils.MDGetValue(ctx, utils.MDKeyUserID)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{}
	userID, err := strconv.ParseUint(mdKeyUserID, 10, 32)
	if err != nil {
		return nil, err
	}
	userModel, err = userModel.FindUserByID(s.DB, uint32(userID))
	if err != nil {
		return nil, err
	}

	userModel.Name = req.Name
	userModel.Password = req.Password
	
	userModel, err = userModel.UpdateUser(s.DB, uint32(userID))
	if err != nil {
		return nil, err
	}

	return &user.UserDetails{
		Id:       int32(userModel.ID),
		Name:     userModel.Name,
		Email:	  userModel.Email,
	}, nil
}