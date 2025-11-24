package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github/http/copy/task2/generated/test"
	"github/http/copy/task2/models"

	"google.golang.org/grpc/codes"
)

type Question struct {
	db *sql.DB
}

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, req *test.CreateQuestionRequest) (*test.CreateQuestionResponse, error)
	GetQuestions(ctx context.Context, req *test.GetQuestionRequest) (*test.GetQuestionResponse, error)
	GetAllQuestions(ctx context.Context, req *test.GetAllQuestionsRequest) (*test.GetAllQuestionsResponse, error)
	SaveResult(ctx context.Context, req *test.SaveResultRequest) (*test.SaveResultResponse, error)
	GetUserAnswersForAttempt(ctx context.Context, req *test.GetUsersAnswerForAttemptRequest) (*test.GetUsersAnswerForAttemptResponse, error)
	GetAttemptList(ctx context.Context, req *test.GetAttemptListRequest) (*test.GetAttemptListResponse, error)
	GetUserHistory(ctx context.Context, req *test.GetUserHistoryRequest) (*test.GetUserHistoryResponse, error)
	CheckAttemptsExist(ctx context.Context, req *test.CheckAttemptsExistRequest) (*test.CheckAttemptsExistResponce, error)
	SaveAttempts(ctx context.Context, req *test.SaveAttemptsRequest) (*test.SaveAttemptsResponse, error)
	GetLastestAttempt(ctx context.Context, req *test.GetLastestAttemptRequest) (*test.GetLastestAttemptResponse, error)
	GetResultByAttempt(ctx context.Context, req *test.GetResultByAttemptRequest) (*test.GetResultByAttemptResponse, error)
	UpdateQuestion(ctx context.Context, req *test.UpdateQuestionRequest) (*test.UpdateQuestionResponse, error)
	DeleteQuestion(ctx context.Context, req *test.DeleteQuestionRequest) (*test.DeleteQuestionResponse, error)
}

func NewQuestion(db *sql.DB) QuestionRepo {
	return &Question{
		db: db,
	}
}

func (p *Question) CreateQuestion(ctx context.Context, req *test.CreateQuestionRequest) (*test.CreateQuestionResponse, error) {

	query := `INSERT INTO questions(name, a, b, true_answer, subject_id, grade_id) 
	VALUES($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(
		query,
		&req.Name,
		&req.A,
		&req.B,
		&req.TrueAnswer,
		&req.SubjectId,
		&req.GradeId,
	)
	if err != nil {
		return nil, err
	}

	return &test.CreateQuestionResponse{
		Message: "success!",
	}, nil
}

func (p *Question) GetQuestions(ctx context.Context, req *test.GetQuestionRequest) (*test.GetQuestionResponse, error) {

	query := `SELECT 
		id, name, a, b, subject_id, grade_id 
		FROM questions
		WHERE subject_id = $1 
		AND grade_id = $2 
		LIMIT 5`

	rows, err := p.db.Query(query, req.SubjectId, req.GradeId)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	resp := &test.GetQuestionResponse{}

	for rows.Next() {
		q := &test.GetQuestionRequest{}
		if err := rows.Scan(
			&q.QuestionId,
			&q.Name,
			&q.A,
			&q.B,
			&q.SubjectId,
			&q.GradeId,
		); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		resp.Questions = append(resp.Questions, q)

		if req.SubjectId != q.SubjectId || req.GradeId != q.GradeId {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
	}

	fmt.Println("-->", resp)

	return resp, nil
}

func (p *Question) SaveResult(ctx context.Context, req *test.SaveResultRequest) (*test.SaveResultResponse, error) {

	if req == nil || len(req.Results) == 0 {
		return nil, errors.New("no results to save")
	}

	query := "INSERT INTO result (user_id, attempt_id, question_id, is_correct, selected_option) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, result := range req.Results {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
			i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		values = append(values, req.UserId, req.LastAttemptId, result.QuestionId, result.IsCorrect, result.SelectedOption)
	}

	if len(placeholders) == 0 {
		return nil, errors.New("no valid results to insert")
	}

	query += strings.Join(placeholders, ",") + ";"

	_, err := p.db.Exec(query, values...)
	if err != nil {
		return nil, err
	}

	return &test.SaveResultResponse{Message: "success!"}, nil
}

func (p *Question) GetAllQuestions(ctx context.Context, req *test.GetAllQuestionsRequest) (*test.GetAllQuestionsResponse, error) {

	rows, err := p.db.Query("SELECT id, name, a, b, true_answer, subject_id, grade_id FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questions := []*test.GetQuestionRequest{}
	for rows.Next() {
		var question models.Questions
		err := rows.Scan(&question.Id, &question.Name, &question.A, &question.B, &question.TruAnsver, &question.SubjectId, &question.GradeId)
		if err != nil {
			return nil, err
		}

		questions = append(questions, &test.GetQuestionRequest{
			QuestionId: question.Id,
			Name:       question.Name,
			A:          question.A,
			B:          question.B,
			TrueAnswer: question.TruAnsver,
			SubjectId:  question.SubjectId,
			GradeId:    question.GradeId,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &test.GetAllQuestionsResponse{
		Questions: questions,
	}, nil
}

func (p *Question) GetUserAnswersForAttempt(ctx context.Context, req *test.GetUsersAnswerForAttemptRequest) (*test.GetUsersAnswerForAttemptResponse, error) {

	query := `
        SELECT user_id, attempt_id, question_id, is_correct, selected_option 
        FROM result 
        WHERE user_id = $1 AND attempt_id = $2
    `
	rows, err := p.db.Query(query, req.UserId, req.AttemptCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result models.Result
		if err := rows.Scan(&result.UserId, &result.AttemptCount, &result.QuestionId, &result.IsCorrect, &result.SelectedOption); err != nil {
			return nil, err
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &test.GetUsersAnswerForAttemptResponse{
		UserId:       req.UserId,
		AttemptCount: req.AttemptCount,
	}, nil
}

func (p *Question) GetAttemptList(ctx context.Context, req *test.GetAttemptListRequest) (*test.GetAttemptListResponse, error) {
	query := `
		SELECT 
			r.attempt_id,
			r.user_id,
			COUNT(CASE WHEN r.is_correct = true THEN 1 END) AS correct_count
		FROM result r
		WHERE r.user_id = $1
		GROUP BY r.attempt_id, r.user_id
		ORDER BY r.attempt_id;
	`

	rows, err := p.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attemptList []*test.GetAttemptListRequest
	for rows.Next() {
		var attempt test.GetAttemptListRequest
		if err := rows.Scan(&attempt.AttemptId, &attempt.UserId, &attempt.CorrectCount); err != nil {
			return nil, err
		}

		attemptList = append(attemptList, &attempt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &test.GetAttemptListResponse{
		AttemptList: attemptList,
	}, nil
}

func (p *Question) GetUserHistory(ctx context.Context, req *test.GetUserHistoryRequest) (*test.GetUserHistoryResponse, error) {

	query := `SELECT attempt_id, question_id, is_correct, user_id, selected_option FROM result WHERE attempt_id=$1 AND user_id=$2`

	rows, err := p.db.Query(query, req.AttemptCount, req.UserId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userHistory []*test.GetUserHistoryRequest
	for rows.Next() {
		r := test.GetUserHistoryRequest{}
		if err := rows.Scan(&r.AttemptCount, &r.QuestionId, &r.IsCorrect, &r.UserId, &r.SelectedOption); err != nil {
			return nil, err
		}

		userHistory = append(userHistory, &r)
	}

	if len(userHistory) == 0 {
		return nil, fmt.Errorf(codes.InvalidArgument.String())
	}

	return &test.GetUserHistoryResponse{
		UserHistory: userHistory,
	}, nil
}

func (p *Question) CheckAttemptsExist(ctx context.Context, req *test.CheckAttemptsExistRequest) (*test.CheckAttemptsExistResponce, error) {
	query := `SELECT COUNT(*) FROM result WHERE attempt_id=$1 AND user_id=$2`

	var count int
	err := p.db.QueryRow(query, req.AttemptCount, req.UserId).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &test.CheckAttemptsExistResponce{
		AttemptCount: req.AttemptCount,
	}, nil
}

func (p *Question) SaveAttempts(ctx context.Context, req *test.SaveAttemptsRequest) (*test.SaveAttemptsResponse, error) {

	query := `INSERT INTO attempts(created_at) VALUES($1)`
	_, err := p.db.Exec(
		query,
		req.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &test.SaveAttemptsResponse{
		Message: "success!",
	}, nil
}

func (p *Question) GetLastestAttempt(ctx context.Context, req *test.GetLastestAttemptRequest) (*test.GetLastestAttemptResponse, error) {

	var attemptCount int

	query := `select attempt_id from result where user_id = $1 order by attempt_id desc limit 1`

	err := p.db.QueryRow(query, req.UserId).Scan(&attemptCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &test.GetLastestAttemptResponse{
		LastAttemptId: int32(attemptCount),
	}, nil
}

func (p *Question) GetResultByAttempt(ctx context.Context, req *test.GetResultByAttemptRequest) (*test.GetResultByAttemptResponse, error) {

	query := `SELECT attempt_id, question_id, is_correct, selected_option FROM result WHERE attempt_id=$1`

	rows, err := p.db.Query(query, req.AttemptId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*test.QuestionResult

	for rows.Next() {
		var a test.QuestionResult
		if err := rows.Scan(&a.AttemptId, &a.QuestionId, &a.IsCorrect, &a.SelectedOption); err != nil {
			return nil, err
		}

		result = append(result, &a)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &test.GetResultByAttemptResponse{
		Results: result,
		UserId:  req.UserId,
	}, nil
}

func (p *Question) UpdateQuestion(ctx context.Context, req *test.UpdateQuestionRequest) (*test.UpdateQuestionResponse, error) {

	query := `Update questions
	SET name=$4,
	a=$5, b=$6,
	true_answer=$7
	WHERE id=$1 AND subject_id=$2 AND grade_id=$3
	`

	_, err := p.db.Exec(
		query,
		req.QuestionId,
		req.SubjectId,
		req.GradeId,
		req.Name,
		req.A,
		req.B,
		req.TrueAnswer,
	)
	if err != nil {
		return nil, err
	}

	return &test.UpdateQuestionResponse{
		QuestionId: req.QuestionId,
		Name:       req.Name,
		A:          req.A,
		B:          req.B,
		TrueAnswer: req.TrueAnswer,
		SubjectId:  req.SubjectId,
		GradeId:    req.GradeId,
	}, nil
}

var ErrNoRows = errors.New("question not found")

func (p *Question) DeleteQuestion(ctx context.Context, req *test.DeleteQuestionRequest) (*test.DeleteQuestionResponse, error) {

	query := `DELETE FROM questions WHERE id=$1`

	result, err := p.db.Exec(query, req.QuestionId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &test.DeleteQuestionResponse{
		Message: "success!",
	}, nil

}
