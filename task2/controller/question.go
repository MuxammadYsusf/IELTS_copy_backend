package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github/http/copy/task2/generated/test"
	"github/http/copy/task2/models"
	"github/http/copy/task2/storage"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

func (c *Handler) CreateQuestion(ctx *gin.Context) {

	var test test.CreateQuestionRequest

	if err := ctx.ShouldBindJSON(&test); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error on binding JSON(C)!",
		})
		return
	}

	resp, err := c.GRPCClient.QuestionService().CreateQuestion(ctx, &test)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on inserting questions!"})
		fmt.Println("-->", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "question inserted successfully!",
		"responce": resp,
	})
}

func (c *Handler) GetQuestion(ctx *gin.Context) {
	var (
		subjectID int
		gradeID   int
	)

	subjectID = cast.ToInt(ctx.Param("subject_id"))
	gradeID = cast.ToInt(ctx.Param("grade_id"))

	rows, err := c.GRPCClient.QuestionService().GetQuestion(ctx, &test.GetQuestionRequest{
		SubjectId: int32(subjectID),
		GradeId:   int32(gradeID),
	})
	if err != nil {
		if rows == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "not found!"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch questions"})
			return
		}
	}

	ctx.JSON(http.StatusOK, rows)
}

func (c *Handler) GetResultByAttempt(ctx *gin.Context) {
	userID := cast.ToInt(ctx.GetInt("user_id"))

	AttemptIDStr := ctx.Param("attempt_id")

	AttemptID, err := strconv.Atoi(AttemptIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ivalid type Id(73)!"})
		return
	}

	exists, err := c.GRPCClient.QuestionService().CheckAttemptsExist(ctx, &test.CheckAttemptsExistRequest{
		AttemptCount: int32(AttemptID),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking attempt existence (82)!"})
		return
	}
	if exists.AttemptCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Attempt not found or does not belong to the user(86)!"})
		return
	}

	resp, err := c.GRPCClient.QuestionService().GetUserHistory(ctx, &test.GetUserHistoryRequest{
		UserId:       int32(userID),
		AttemptCount: int32(AttemptID),
	})
	if err != nil {
		if resp == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "not found!"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on getting user history by attempt(114)!"})
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)

}

func (c *Handler) GetAttemptsList(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	resp, err := c.GRPCClient.QuestionService().GetAttemptList(ctx, &test.GetAttemptListRequest{
		UserId: int32(userID),
	})
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on getting user history by attempt(114)!"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Handler) AnswerToQuestions(ctx *gin.Context) {
	var request struct {
		Items []*test.QuestionResult `json:"items"`
	}

	userID := ctx.GetInt("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при привязке JSON!"})
		return
	}

	resp, err := c.GRPCClient.QuestionService().CheckAnswers(ctx, &test.CheckAnswersRequest{
		UserId: int32(userID),
		Items:  request.Items,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on testing process!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Тест завершен",
		"result":  resp,
	})
}

func (c *Handler) UpdateQuestion(ctx *gin.Context) {
	var q models.Questions
	if err := ctx.ShouldBindJSON(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error on binding JSON(U)!"})
		return
	}

	_, err := c.GRPCClient.QuestionService().UpdateQuestion(ctx, &test.UpdateQuestionRequest{
		QuestionId: int32(q.Id),
		Name:       q.Name,
		TrueAnswer: q.TruAnsver,
		SubjectId:  int32(q.SubjectId),
		GradeId:    int32(q.GradeId),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on updating questions!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "question updated successfully!"})
}

func (c *Handler) DeleteQuestion(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ivalid type Id!"})
		return
	}

	_, err = c.GRPCClient.QuestionService().DeleteQuestion(ctx, &test.DeleteQuestionRequest{
		QuestionId: int32(id),
	})
	if err == storage.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no rows affected!"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on deleting question!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "question deleted successfully!"})
}
