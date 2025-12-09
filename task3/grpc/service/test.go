package service

import (
	"context"
	"errors"
	"fmt"
	"github/http/copy/task3/generated/test"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func getBandScore(correctAnswers int) float64 {
	switch {
	case correctAnswers >= 39:
		return 9.0
	case correctAnswers >= 37:
		return 8.5
	case correctAnswers >= 35:
		return 8.0
	case correctAnswers >= 33:
		return 7.5
	case correctAnswers >= 30:
		return 7.0
	case correctAnswers >= 27:
		return 6.5
	case correctAnswers >= 23:
		return 6.0
	case correctAnswers >= 19:
		return 5.5
	case correctAnswers >= 15:
		return 5.0
	case correctAnswers >= 13:
		return 4.5
	case correctAnswers >= 10:
		return 4.0
	case correctAnswers >= 5:
		return 3.5
	case correctAnswers >= 3:
		return 3.5
	default:
		return 1
	}
}

func (s *TestService) NewTest(ctx context.Context, req *test.NewTestRequest) (*test.NewTestResponse, error) {

	resp, err := s.testPostgres.Test().CreateNewTest(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) Writing(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error) {

	resp, err := s.testPostgres.Test().CreateWritingQuestions(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) GetWritingQuestions(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error) {

	resp, err := s.testPostgres.Test().GetWritingQuestions(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) Speaking(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error) {

	resp, err := s.testPostgres.Test().CreateSpeakingQuestions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TestService) GetSpeakingQuestions(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error) {

	resp, err := s.testPostgres.Test().GetSpeakingQuestions(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) ReadingQuestions(ctx context.Context, req *test.ReadingQuestionRequest) (*test.ReadingResponse, error) {

	resp, err := s.testPostgres.Test().ReadingQuestions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TestService) ReadingContent(ctx context.Context, req *test.ReadingContentRequest) (*test.ReadingResponse, error) {

	resp, err := s.testPostgres.Test().ReadingContent(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TestService) GetReadingPassages(ctx context.Context, req *test.GetReadingPessagesRequest) (*test.GetReadingPessagesResponse, error) {

	resp, err := s.testPostgres.Test().GetReadingPassages(ctx, req)
	if err != nil {
		return nil, err
	}

	lastestAttempt, err := s.testPostgres.Test().RGetLastAttemptId(ctx, &test.RGetLastestAttemptRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	lastestAttemptCount := lastestAttempt.LastAttemptId + 1

	start := time.Now()
	var end time.Time

	if len(resp.Bodies) == 1 {
		end = start.Add(time.Minute * 20)
	} else if len(resp.Bodies) == 2 {
		end = start.Add(time.Minute * 40)
	} else if len(resp.Bodies) == 3 {
		end = start.Add(time.Hour)
	}

	if _, err = s.testPostgres.Test().SaveReadingAttempts(ctx, &test.ReadingAttemptsRequest{
		ReadingAtId: lastestAttemptCount,
		UserId:      req.UserId,
		Start:       timestamppb.New(start),
		End:         timestamppb.New(end),
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func normalizeSet(list *test.StringList) map[string]struct{} {

	out := make(map[string]struct{}, len(list.Values))
	for _, v := range list.Values {
		cleaned := strings.TrimSpace(strings.ToLower(v))
		out[cleaned] = struct{}{}
	}

	return out
}

func (s *TestService) CheckReadingAnswers(ctx context.Context, req *test.CheckReadingAnswersRequest) (*test.CheckReadingAnswersResponse, error) {

	var orderIds []int32
	var passageIds []int32
	var checkedAnswers []*test.ReadingAnswers

	for i := range req.Passages {
		var passageId int32
		if req.Passages[i].OrderId <= 13 {
			passageId = 1
		} else if req.Passages[i].OrderId <= 26 {
			passageId = 2
		} else if req.Passages[i].OrderId <= 40 {
			passageId = 3
		}

		checked := &test.ReadingAnswers{
			PassageId: passageId,
			OrderId:   req.Passages[i].OrderId,
			Answer:    req.Passages[i].Answer,
		}

		checkedAnswers = append(checkedAnswers, checked)

	}

	for _, ans := range checkedAnswers {
		passageIds = append(passageIds, ans.PassageId)
	}

	for _, item := range req.Passages {
		orderIds = append(orderIds, item.OrderId)
	}

	resp, err := s.testPostgres.Test().GetReadingTrueAnswers(ctx, &test.GetReadingTrueAnswersBatchRequest{
		OrderIds:   orderIds,
		PassageIds: passageIds,
		TestId:     req.TestId,
	})
	if err != nil {
		return nil, err
	}

	// ???
	timeResp, err := s.testPostgres.Test().GetReadingRealTime(ctx, &test.ReadingAttemptsRequest{
		UserId:      req.UserId,
		ReadingAtId: req.AttemptId,
	})
	if err != nil {
		return nil, err
	}
	bufferTime := time.Second * 2
	if time.Now().After(timeResp.End.AsTime().Add(bufferTime)) {
		return nil, errors.New("the time is out")
	}

	if req.AttemptId == 0 {
		return nil, errors.New("lastestAttempt is nil — no record found in r_attempts")
	}

	lastestAttemptCount := req.AttemptId
	results := req.Passages
	correctCount := int32(0)

	fmt.Println(lastestAttemptCount)

	trueAnswerMap := make(map[int32]map[string]struct{}, len(resp.TrueAnswers))
	for orderID := range resp.TrueAnswers {
		trueAnswerMap[orderID] = normalizeSet(resp.TrueAnswers[orderID])
	}

	for i, ans := range results {
		normolized := strings.TrimSpace(strings.ToLower(ans.Answer))

		if correctSet, ok := trueAnswerMap[int32(results[i].OrderId)]; ok {
			if _, found := correctSet[normolized]; found {
				results[i].IsCorrect = true
				correctCount++
			}
		}
	}

	score := getBandScore(int(correctCount))

	var grpcResult []*test.ReadingResultRequest

	for i, item := range results {

		orderID := item.OrderId
		userAnswer := item.Answer

		grpcResult = append(grpcResult, &test.ReadingResultRequest{
			OrderId:   orderID,
			Answer:    userAnswer,
			IsCorrect: results[i].IsCorrect,
		})
	}
	testId := req.TestId

	_, err = s.testPostgres.Test().SaveReadingResult(ctx, &test.ReadingResultRequest{
		UserId:    req.UserId,
		AttemptId: lastestAttemptCount,
		Results:   grpcResult,
		Score:     float32(score),
		TestId:    testId,
	})
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	response := &test.CheckReadingAnswersResponse{
		UserId:       req.UserId,
		Score:        float32(score),
		CorrectCount: correctCount,
		Result:       results,
	}

	return response, nil
}

func (s *TestService) ListeningQuestions(ctx context.Context, req *test.ListeningQuestionRequest) (*test.ListeningResponse, error) {

	resp, err := s.testPostgres.Test().CreateListeningQuestions(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) ListeningContent(ctx context.Context, req *test.ListeningContentRequest) (*test.ListeningResponse, error) {

	resp, err := s.testPostgres.Test().CreateListeningContent(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TestService) GetListeningContent(ctx context.Context, req *test.GetListeningContentRequest) (*test.GetListeningContentResponse, error) {

	resp, err := s.testPostgres.Test().GetListeningContent(ctx, req)
	if err != nil {
		return nil, err
	}

	laastestAttempt, err := s.testPostgres.Test().LGetLastestAttempt(ctx, &test.LLastestAttemptRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	lastestAttemptCount := laastestAttempt.LastAttemptId + 1

	start := time.Now()
	var end time.Time

	if len(resp.Bodies) == 1 {
		end = start.Add(time.Minute * 8)
	} else if len(resp.Bodies) == 2 {
		end = start.Add(time.Minute * 16)
	} else if len(resp.Bodies) == 3 {
		end = start.Add(time.Minute * 24)
	} else if len(resp.Bodies) == 4 {
		end = start.Add(time.Minute * 32)
	}

	_, err = s.testPostgres.Test().SaveListeningAttempts(ctx, &test.ListeningAttemptsRequest{
		ListeningAtId: lastestAttemptCount,
		UserId:        req.UserId,
		Start:         timestamppb.New(start),
		End:           timestamppb.New(end),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TestService) CheckListeningAnswers(ctx context.Context, req *test.CheckListeningAnswersRequest) (*test.CheckListeningAnswersResponse, error) {

	var orderIds []int32
	var sectionIds []int32
	var checkedAnswers []*test.ListeningAnswer

	for _, item := range req.Sections {
		orderIds = append(orderIds, item.OrderId)
	}
	for i := range req.Sections {
		var sectionID int32
		if req.Sections[i].SectionId <= 10 {
			sectionID = 1
		} else if req.Sections[i].SectionId <= 20 {
			sectionID = 2
		} else if req.Sections[i].SectionId <= 30 {
			sectionID = 3
		} else if req.Sections[i].SectionId <= 40 {
			sectionID = 4
		}

		checked := &test.ListeningAnswer{
			SectionId: sectionID,
			OrderId:   req.Sections[i].OrderId,
			Answer:    req.Sections[i].Answer,
		}

		checkedAnswers = append(checkedAnswers, checked)

	}

	for _, ans := range checkedAnswers {
		sectionIds = append(sectionIds, ans.SectionId)
	}

	resp, err := s.testPostgres.Test().GetListeningTrueAnswers(ctx, &test.GetListeningTrueAnswersBatchRequest{
		OrderIds:   orderIds,
		SectionIds: sectionIds,
	})
	if err != nil {
		return nil, err
	}

	timeResp, err := s.testPostgres.Test().GetListeningRealTime(ctx, &test.ListeningAttemptsRequest{
		UserId: req.UserId,
	},
	)
	if err != nil {
		return nil, err
	}
	bufferTime := time.Second * 2
	if time.Now().After(timeResp.End.AsTime().Add(bufferTime)) {
		return nil, errors.New("the time is out")
	}

	if req.AttemptId == 0 {
		return nil, errors.New("lastestAttempt is nil — no record found in r_attempts")
	}

	lastestAttemptCount := req.AttemptId
	correctCount := int32(0)
	result := req.Sections

	trueAmswerMap := make(map[int32]map[string]struct{}, len(resp.TrueAnswers))
	for orderID := range resp.TrueAnswers {
		trueAmswerMap[orderID] = normalizeSet(resp.TrueAnswers[orderID])
	}

	for i, ans := range result {
		normolized := strings.TrimSpace(strings.ToLower(ans.Answer))
		if correctSet, ok := trueAmswerMap[result[i].OrderId]; ok {
			if _, found := correctSet[normolized]; found {
				result[i].IsCorrect = true
				correctCount++
			}
		}
	}

	score := getBandScore(int(correctCount))

	var grpcResult []*test.ListeningResultRequest

	for i, item := range result {

		orderID := item.OrderId
		Answers := item.Answer

		testId := req.TestId

		grpcResult = append(grpcResult, &test.ListeningResultRequest{
			OrderId:   orderID,
			Answer:    Answers,
			IsCorrect: result[i].IsCorrect,
			TestId:    testId,
		})
	}

	_, err = s.testPostgres.Test().SaveListeningAnswers(ctx, &test.ListeningResultRequest{
		UserId:        req.UserId,
		ListeningAtId: lastestAttemptCount,
		Results:       grpcResult,
		Score:         float32(score),
	})
	if err != nil {
		return nil, err
	}

	response := &test.CheckListeningAnswersResponse{
		UserId:       req.UserId,
		Score:        float32(score),
		Result:       result,
		CorrectCount: correctCount,
	}

	return response, nil
}
