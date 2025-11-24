package service

import (
	"context"
	"errors"
	"github/http/copy/task2/generated/test"
	"github/http/copy/task2/storage"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *QuestionService) CreateQuestion(ctx context.Context, req *test.CreateQuestionRequest) (*test.CreateQuestionResponse, error) {

	resp, err := s.Questionstorage.Question().CreateQuestion(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, req *test.GetQuestionRequest) (*test.GetQuestionResponse, error) {

	resp, err := s.Questionstorage.Question().GetQuestions(ctx, req)

	if err != nil {
		return nil, err
	}
	if len(resp.Questions) == 0 {
		return nil, errors.New(codes.InvalidArgument.String())
	}

	return resp, nil
}

func (s *QuestionService) GetAllQuestions(ctx context.Context, req *test.GetAllQuestionsRequest) (*test.GetAllQuestionsResponse, error) {

	QuestionResponse, err := s.Questionstorage.Question().GetAllQuestions(ctx, req)
	if err != nil {
		return nil, err
	}

	return QuestionResponse, nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, req *test.UpdateQuestionRequest) (*test.UpdateQuestionResponse, error) {

	resp, err := s.Questionstorage.Question().UpdateQuestion(ctx, req)
	if err != nil {
		return nil, nil
	}

	return resp, nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, req *test.DeleteQuestionRequest) (*test.DeleteQuestionResponse, error) {

	resp, err := s.Questionstorage.Question().DeleteQuestion(ctx, req)
	if resp == nil {
		return nil, storage.ErrNoRows
	} else if err != nil {
		return nil, err
	}

	return resp, nil
}

type SaveResultRequest struct {
	Req     *test.SaveResultRequest
	Results []*test.SaveResultRequest
}

func (s *QuestionService) SaveResult(ctx context.Context, req *test.SaveResultRequest) (*test.SaveResultResponse, error) {

	resp, err := s.Questionstorage.Question().SaveResult(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *QuestionService) GetAttemptList(ctx context.Context, req *test.GetAttemptListRequest) (*test.GetAttemptListResponse, error) {

	resp, err := s.Questionstorage.Question().GetAttemptList(ctx, req)
	if err != nil {
		return nil, nil
	}

	return resp, nil
}

func (s *QuestionService) GetUserHistory(ctx context.Context, req *test.GetUserHistoryRequest) (*test.GetUserHistoryResponse, error) {

	resp, err := s.Questionstorage.Question().GetUserHistory(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resp.UserHistory) == 0 {
		return nil, errors.New(codes.InvalidArgument.String())
	}

	return resp, nil
}

func (s *QuestionService) CheckAttemptsExist(ctx context.Context, req *test.CheckAttemptsExistRequest) (*test.CheckAttemptsExistResponce, error) {

	resp, err := s.Questionstorage.Question().CheckAttemptsExist(ctx, req)
	if err != nil {
		return nil, nil
	}

	return resp, nil
}

func (s *QuestionService) CheckAnswers(ctx context.Context, req *test.CheckAnswersRequest) (*test.CheckAnswersResponse, error) {
	if req == nil || len(req.Items) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Нет данных для проверки")
	}

	questionResponse, err := s.Questionstorage.Question().GetAllQuestions(ctx, &test.GetAllQuestionsRequest{})
	if err != nil {
		return nil, err
	}

	latestAttemptResponse, err := s.Questionstorage.Question().GetLastestAttempt(ctx, &test.GetLastestAttemptRequest{
		UserId: req.UserId,
	})
	if err != nil || questionResponse == nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при получении вопросов: %v", err)
	}
	if err != nil || latestAttemptResponse == nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при получении последней попытки: %v", err)
	}

	latestAttemptCount := latestAttemptResponse.LastAttemptId + 1
	inCorrectCount := 0
	questions := questionResponse.Questions
	results := req.Items

	for i := range req.Items {
		for _, question := range questions {
			if int32(req.Items[i].QuestionId) == question.QuestionId {
				req.Items[i].IsCorrect = req.Items[i].SelectedOption == question.TrueAnswer
				if !req.Items[i].IsCorrect {
					inCorrectCount++
				}
				break
			}
		}
	}

	_, err = s.Questionstorage.Question().SaveAttempts(ctx, &test.SaveAttemptsRequest{
		AttemptCount: int32(latestAttemptCount),
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return nil, err
	}

	var grpcResults []*test.QuestionResult
	for i := range results {
		grpcResults = append(grpcResults, &test.QuestionResult{
			QuestionId:     results[i].QuestionId,
			SelectedOption: results[i].SelectedOption,
			IsCorrect:      results[i].IsCorrect,
		})
	}

	_, err = s.Questionstorage.Question().SaveResult(ctx, &test.SaveResultRequest{
		LastAttemptId: int32(latestAttemptCount),
		UserId:        int32(req.UserId),
		Results:       grpcResults,
	})
	if err != nil {
		return nil, err
	}

	response := &test.CheckAnswersResponse{
		UserId:         req.UserId,
		AttemptCount:   int32(latestAttemptCount),
		Results:        req.Items,
		IncorrectCount: int32(inCorrectCount),
	}

	return response, nil
}

func (s *QuestionService) SaveAttempts(ctx context.Context, req *test.SaveAttemptsRequest) (*test.SaveAttemptsResponse, error) {

	resp, err := s.Questionstorage.Question().SaveAttempts(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *QuestionService) GetResulByAttempt(ctx context.Context, req *test.GetResultByAttemptRequest) (*test.GetResultByAttemptResponse, error) {

	resp, err := s.Questionstorage.Question().GetResultByAttempt(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
