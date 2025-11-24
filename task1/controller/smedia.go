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

func (c *Controler) CreateSmedia(ctx *gin.Context) {
	var smedia models.Models

	org_author_id := ctx.GetHeader("author_id_from_token")

	if err := ctx.ShouldBindJSON(&smedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	num, err := strconv.Atoi(org_author_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid author id"})
	}

	if smedia.Author_id != num {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	response, err := storage.GetAllSmedia(c.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't get smedia by id!"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controler) GetAllSmedia(ctx *gin.Context) {
	if _, err := storage.GetAllSmedia(c.db); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't get all smedia!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully got all media!"})

}

func (c *Controler) GetByIdSmedia(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}

	response, err := storage.GetByIdSmedia(c.db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't get smedia by id!"})
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (c *Controler) UpdateSmedia(ctx *gin.Context) {
	var smedia models.Models

	org_author_id := ctx.GetHeader("author_id_from_token")

	if err := ctx.ShouldBindJSON(&smedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	num, err := strconv.Atoi(org_author_id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "successfully got all media!"})
		return
	}
	if smedia.Author_id != num {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "permission denied!"})
		return
	}

	if _, err := storage.UpdateSmedia(c.db, smedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error on updating media!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully got all media!"})
}

func (c *Controler) DeleteSmedia(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{"error": "bad requests!"})
		return
	}

	err = storage.DeleteSmedia(c.db, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			ctx.JSON(http.StatusOK, gin.H{"error": "media not found!"})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "couldn't delete media!"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully got all media!"})

}
