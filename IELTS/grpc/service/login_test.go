package service_test

import (
	"context"
	"github/http/copy/task3/generated/session"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginValidation(t *testing.T) {

	tests := []struct {
		name      string
		payload   *session.LoginRequest
		wantSatus int
		wantError string
	}{
		{
			name: "invalid name",
			payload: &session.LoginRequest{
				Name:     "№@mE#",
				Password: "password",
			},
			wantSatus: http.StatusBadRequest,
			wantError: "invalid name or password",
		},
		{
			name: "tricky password",
			payload: &session.LoginRequest{
				Name:     "name",
				Password: "SELECT * FROM table",
			},
			wantSatus: http.StatusBadRequest,
			wantError: "invalid name or password",
		},
		{
			name: "no name",
			payload: &session.LoginRequest{
				Name:     "",
				Password: "password",
			},
			wantSatus: http.StatusBadRequest,
			wantError: "invalid name or password",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			resp, err := loginService.Login(context.Background(), v.payload)

			assert.Nil(t, resp)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), v.wantError, v.wantSatus)
		})
	}
}
