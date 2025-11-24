package controller

import (
	"fmt"
	"net/http"

	"github/http/copy/task1/models"
	"github/http/copy/task1/storage"

	"github.com/gin-gonic/gin"
)

func (c *Controler) LikeHandler(ctx *gin.Context) {

	var likeReq models.LikeRequest
	if err := ctx.ShouldBindJSON(likeReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding JSON!"})
		return
	}
	exists, currentLiked, err := storage.CheckLikeExists(c.db, likeReq.AuthorID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "loke does not exist!"})
		return
	}

	if exists {
		newLiked := !currentLiked
		if newLiked {
			err = storage.AddLike(c.db, likeReq)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on inserting like!"})
				return
			}
			ctx.JSON(http.StatusOK, likeReq)
		} else {
			err = storage.RemoveLike(c.db, likeReq)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on inserting like!"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "like removed!"})
		}
	} else {
		err = storage.AddLike(c.db, likeReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on inserting like!"})
			return
		}
		fmt.Println("like added!")
		ctx.JSON(http.StatusOK, likeReq)
	}
}

func (c *Controler) CountLikesHandler(ctx *gin.Context) {

	postID := ctx.Query("id")
	if postID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Необходимо указать параметр id!"})
		return
	}

	count, err := storage.CountLikes(c.db, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подсчете лайков!"})
		return
	}

	response := map[string]interface{}{
		"id":         postID,
		"like_count": count,
	}

	ctx.JSON(http.StatusOK, response)

}
