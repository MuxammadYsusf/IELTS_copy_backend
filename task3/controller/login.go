package controller

import (
	"fmt"
	"github/http/copy/task3/generated/session"
	"github/http/copy/task3/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Reg(ctx *gin.Context) {
	var reg session.RegisterRequest

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.LoginService().Register(ctx, &reg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) Login(ctx *gin.Context) {
	var l session.LoginRequest

	if err := ctx.ShouldBindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.LoginService().Login(ctx, &l)
	if err != nil {
		if fmt.Sprintf("%v", err) == "rpc error: code = Unknown desc = user not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenStr, err := security.GenerateJWTToken(int(resp.UserId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}
