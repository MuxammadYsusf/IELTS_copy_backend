package controller

import (
	"fmt"
	"github/http/copy/task3/generated/test"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (c *Controller) CreateNewTest(ctx *gin.Context) {

	var t test.NewTestRequest

	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().NewTest(ctx, &t)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateWritingQuestions(ctx *gin.Context) {

	var w test.WritingRequest

	if err := ctx.ShouldBindJSON(&w); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().Writing(ctx, &w)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetWritingQuestions(ctx *gin.Context) {
	var testId int
	var taskId int

	testId = cast.ToInt(ctx.Param("testId"))
	taskId = cast.ToInt(ctx.Param("taskId"))

	resp, err := c.GRPCClient.TestService().GetWritingQuestions(ctx, &test.WritingRequest{
		TestId: int32(testId),
		TaskId: int32(taskId),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateSpeakingQuestions(ctx *gin.Context) {

	var s test.SpeakingRequest

	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().Speaking(ctx, &s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetSpeakingQuestions(ctx *gin.Context) {
	var testId int
	var partId int

	testId = cast.ToInt(ctx.Param("testId"))
	partId = cast.ToInt(ctx.Param("partId"))

	resp, err := c.GRPCClient.TestService().GetSpeakingQuestions(ctx, &test.SpeakingRequest{
		TestId: int32(testId),
		PartId: int32(partId),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateReadingQuestions(ctx *gin.Context) {

	var r test.ReadingQuestionRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().ReadingQuestions(ctx, &r)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateReadinContent(ctx *gin.Context) {

	var r test.ReadingContentRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().ReadingContent(ctx, &r)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetReadinContent(ctx *gin.Context) {

	var content test.GetReadingPessagesRequest

	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().GetReadingPassages(ctx, &content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) SaveReadingAnswers(ctx *gin.Context) {

	var req struct {
		Passages  []*test.ReadingAnswers `json:"passages"`
		TestId    int32                  `json:"testId"`
		AttemptId int32                  `json:"attemptId"`
	}
	userId := ctx.GetInt("userId")
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().CheckReadingAnswers(ctx, &test.CheckReadingAnswersRequest{
		UserId:    int32(userId),
		Passages:  req.Passages,
		TestId:    req.TestId,
		AttemptId: req.AttemptId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	correctCount := resp.CorrectCount

	ctx.JSON(http.StatusOK, &test.ReadingResponse{
		Message:      "test is finished",
		Score:        resp.Score,
		CorrectCount: correctCount,
		Results:      resp.Result,
	})
}

func (c *Controller) CreateListeningQuestions(ctx *gin.Context) {

	var l test.ListeningQuestionRequest

	if err := ctx.ShouldBindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().ListeningQuestions(ctx, &l)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateListeningContent(ctx *gin.Context) {

	var l test.ListeningContentRequest

	if err := ctx.ShouldBindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().ListeningContent(ctx, &l)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetListeningContent(ctx *gin.Context) {

	var l test.GetListeningContentRequest

	if err := ctx.ShouldBindJSON(&l); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.GRPCClient.TestService().GetListeningContent(ctx, &l)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) SaveListeningAnswers(ctx *gin.Context) {

	var req struct {
		Sections  []*test.ListeningAnswer `json:"sections"`
		TestId    int32                   `json:"testId"`
		AttemptId int32                   `json:"attemptId"`
	}

	userId := ctx.GetInt("userId")
	if userId == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("testId ->", req.TestId)

	resp, err := c.GRPCClient.TestService().CheckListeningAnswers(ctx, &test.CheckListeningAnswersRequest{
		UserId:    int32(userId),
		Sections:  req.Sections,
		TestId:    req.TestId,
		AttemptId: req.AttemptId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	correctCount := resp.CorrectCount

	ctx.JSON(http.StatusOK, &test.ListeningResponse{
		Message:      "test is finished",
		Score:        resp.Score,
		CorrectCount: correctCount,
		Result:       resp.Result,
	})

}
