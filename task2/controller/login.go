package controller

import (
	"github/http/copy/task2/generated/session"
	"github/http/copy/task2/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Handler) LogIn(ctx *gin.Context) {

	var u session.LoginRequest

	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding JSON(L)!"})
		return
	}

	resp, err := c.GRPCClient.LoginService().Login(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "login failed!"})
		return
	}

	tokenStr, err := security.CreateJWT(int(resp.UserId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on generating token!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (c *Handler) Register(ctx *gin.Context) {

	var reg session.RegisterRequest

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding JSON!"})
		return
	}

	if reg.Name == "" || reg.Password == "" || reg.PhoneNumber == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no rows!"})
		return
	}

	resp, err := c.GRPCClient.LoginService().Register(ctx, &reg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "reg failed!"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Handler) UpdatePassword(ctx *gin.Context) {
	var creds session.UpdatePasswordRequest

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding JSON(G)!"})
		return
	}

	resp, err := c.GRPCClient.LoginService().UpdatePassword(ctx, &creds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on updating password!"})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
