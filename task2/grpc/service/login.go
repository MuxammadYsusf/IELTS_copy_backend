package service

import (
	"context"
	"github/http/copy/task2/generated/session"
)

func (s *LoginService) Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error) {

	resp, err := s.Loginstorage.Login().Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LoginService) Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error) {

	resp, err := s.Loginstorage.Login().Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LoginService) UpdatePassword(ctx context.Context, req *session.UpdatePasswordRequest) (*session.UpdatePasswordResponse, error) {

	resp, err := s.Loginstorage.Login().UpdatePassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
