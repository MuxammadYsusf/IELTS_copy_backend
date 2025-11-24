package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github/http/copy/task1/models"
	"github/http/copy/task1/storage"

	"github.com/gin-gonic/gin"
)

func (c *Controler) Reg(ctx *gin.Context) {
	var reg models.LoginRequest
	// Прямое связывание JSON в структуру
	if err := ctx.ShouldBindJSON(&reg); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Вставка данных в хранилище
	if err := storage.InsertReg(c.db, reg); err != nil {
		fmt.Println("Error inserting registration:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully!"})

}

func (c *Controler) DeleteLogInfo(ctx *gin.Context) {

	idStr := ctx.Param("id")
	fmt.Println("id:", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = storage.DeleteLog(c.db, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
			return
		} else {
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted!"})

}
