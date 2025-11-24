package service

import (
	"context"
	"fmt"
	"github/http/copy/task3/generated/session"
	"strings"
)

func (s *LoginService) Register(ctx context.Context, req *session.RegisterRequest) (*session.RegisterResponse, error) {

	if req.Name == "" || req.PhoneNumber == 0 || req.Password == "" {
		return nil, fmt.Errorf("invalid name or password or phone number")
	}

	if strings.ContainsAny(req.Name, "!@#$%&*()№:;<>?'\";--") {
		return nil, fmt.Errorf("invalid name")
	}

	if strings.ContainsAny(req.Password, "'\";--") {
		return nil, fmt.Errorf("invalid password")
	}

	if len(req.Name) > 15 || len(req.Password) > 50 {
		return nil, fmt.Errorf("invalid name or password")
	}

	resp, err := s.loginPostgres.Login().Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (s *LoginService) Login(ctx context.Context, req *session.LoginRequest) (*session.LoginResponse, error) {

	if req.Name == "" || req.Password == "" {
		return nil, fmt.Errorf("invalid name or password")
	}

	if strings.ContainsAny(req.Name, "!@#$%&*()№:;<>?") {
		return nil, fmt.Errorf("invalid name or password")
	}

	if strings.ContainsAny(req.Password, "'\";--") {
		return nil, fmt.Errorf("invalid name or password")
	}

	if len(req.Name) > 15 || len(req.Password) > 50 {
		return nil, fmt.Errorf("invalid name or password")
	}

	resp, err := s.loginPostgres.Login().Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
