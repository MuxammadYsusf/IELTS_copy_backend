package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github/http/copy/task3/generated/test"
	"time"

	"strings"

	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type tests struct {
	db *sql.DB
}

type TestRepo interface {
	CreateNewTest(ctx context.Context, req *test.NewTestRequest) (*test.NewTestResponse, error)
	CreateWritingQuestions(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error)
	GetWritingQuestions(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error)
	CreateSpeakingQuestions(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error)
	GetSpeakingQuestions(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error)
	ReadingQuestions(ctx context.Context, req *test.ReadingQuestionRequest) (*test.ReadingResponse, error)
	ReadingContent(ctx context.Context, req *test.ReadingContentRequest) (*test.ReadingResponse, error)
	GetReadingTrueAnswers(ctx context.Context, req *test.GetReadingTrueAnswersBatchRequest) (*test.GetReadingTrueAnswersBatchResponse, error)
	GetReadingPassages(ctx context.Context, req *test.GetReadingPessagesRequest) (*test.GetReadingPessagesResponse, error)
	SaveReadingResult(ctx context.Context, req *test.ReadingResultRequest) (*test.ReadingResponse, error)
	GetReadingRealTime(ctx context.Context, req *test.ReadingAttemptsRequest) (*test.ReadingAttemptsResponse, error)
	RGetLastAttemptId(ctx context.Context, req *test.RGetLastestAttemptRequest) (*test.RGetLastestAttemptResponse, error)
	SaveReadingAttempts(ctx context.Context, req *test.ReadingAttemptsRequest) (*test.ReadingAttemptsResponse, error)
	CreateListeningQuestions(ctx context.Context, req *test.ListeningQuestionRequest) (*test.ListeningResponse, error)
	CreateListeningContent(ctx context.Context, req *test.ListeningContentRequest) (*test.ListeningResponse, error)
	GetListeningContent(ctx context.Context, req *test.GetListeningContentRequest) (*test.GetListeningContentResponse, error)
	GetListeningTrueAnswers(ctx context.Context, req *test.GetListeningTrueAnswersBatchRequest) (*test.GetListeningTrueAnswersBatchResponse, error)
	GetListeningRealTime(ctx context.Context, req *test.ListeningAttemptsRequest) (*test.ListeningAttemptsResponse, error)
	SaveListeningAnswers(ctx context.Context, req *test.ListeningResultRequest) (*test.ListeningResponse, error)
	SaveListeningAttempts(ctx context.Context, req *test.ListeningAttemptsRequest) (*test.ListeningResponse, error)
	LGetLastestAttempt(ctx context.Context, req *test.LLastestAttemptRequest) (*test.LLastestAttemptResponse, error)
}

func NewTest(db *sql.DB) TestRepo {
	return &tests{
		db: db,
	}
}

func (t *tests) CreateNewTest(ctx context.Context, req *test.NewTestRequest) (*test.NewTestResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO test(name) VALUES($1)`

	_, err = tx.Exec(
		query,
		req.Name,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.NewTestResponse{
		Message: "success",
	}, nil
}

func (t *tests) CreateWritingQuestions(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO writing(task_id, test_id, question) VALUES($1, $2, $3)`

	_, err = tx.Exec(
		query,
		req.TaskId,
		req.TestId,
		req.Question,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.WritingResponse{
		Message: "success",
	}, nil
}

func (t *tests) GetWritingQuestions(ctx context.Context, req *test.WritingRequest) (*test.WritingResponse, error) {

	query := `SELECT question FROM writing WHERE task_id = $1 AND test_id = $2`

	err := t.db.QueryRow(query, req.TaskId, req.TestId).Scan(&req.Question)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("question not found")
		}
		return nil, err
	}

	question := req.Question

	return &test.WritingResponse{
		Question: question,
	}, nil

}

func (t *tests) CreateSpeakingQuestions(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO speaking(part_id, test_id, question) VALUES($1, $2, $3)`

	_, err = tx.Exec(
		query,
		req.PartId,
		req.TestId,
		req.Question,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.SpeakingResponse{
		Message: "success",
	}, nil
}

func (t *tests) GetSpeakingQuestions(ctx context.Context, req *test.SpeakingRequest) (*test.SpeakingResponse, error) {

	err := t.db.QueryRow(`SELECT question FROM speaking WHERE part_id = $1 AND test_id = $2`, req.PartId, req.TestId).Scan(&req.Question)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("question not found")
		}
		return nil, err
	}

	question := req.Question

	return &test.SpeakingResponse{
		Question: question,
	}, nil
}

func (t *tests) ReadingContent(ctx context.Context, req *test.ReadingContentRequest) (*test.ReadingResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO r_contents(passage_id, body, test_id) VALUES($1, $2, $3)`

	_, err = tx.Exec(
		query,
		req.PassageId,
		req.Body,
		req.TestId,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ReadingResponse{
		Message: "success",
	}, nil
}

func (t *tests) ReadingQuestions(ctx context.Context, req *test.ReadingQuestionRequest) (*test.ReadingResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO r_questions(pessage_id, order_id, true_answers, content_id, test_id) VALUES($1, $2, $3, $4, $5)`

	jsonAnswers, err := json.Marshal(req.TrueAnswers)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		query,
		req.PessageId,
		req.OrderId,
		string(jsonAnswers),
		req.ContentId,
		req.TestId,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ReadingResponse{
		Message: "success",
	}, nil
}

func (t *tests) GetReadingPassages(ctx context.Context, req *test.GetReadingPessagesRequest) (*test.GetReadingPessagesResponse, error) {

	PBodies := make(map[string]string)

	for _, id := range req.PassageIds {
		var body string
		err := t.db.QueryRow(`SELECT body FROM r_contents WHERE passage_id = $1 AND test_id = $2`,
			id, req.TestId).Scan(&body)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			} else if id > 3 {
				return nil, err
			}
			return nil, err
		}
		PBodies[fmt.Sprintf("passage%d", id)] = body
	}

	return &test.GetReadingPessagesResponse{
		Bodies: PBodies,
	}, nil
}

func (t *tests) GetReadingTrueAnswers(ctx context.Context, req *test.GetReadingTrueAnswersBatchRequest) (*test.GetReadingTrueAnswersBatchResponse, error) {

	args := make([]interface{}, len(req.OrderIds))
	for i, v := range req.OrderIds {
		args[i] = v
	}

	query := `SELECT order_id,true_answers
			FROM r_questions
			WHERE order_id = ANY($1) 
			AND pessage_id = ANY($2) AND test_id = $3`

	rows, err := t.db.Query(query, pq.Array(req.OrderIds), pq.Array(req.PassageIds), req.TestId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make(map[int32]*test.StringList)

	for rows.Next() {
		var order int32
		var answers []byte

		if err := rows.Scan(&order, &answers); err != nil {
			return nil, err
		}

		var raw map[string]struct {
			Values []string `json:"values"`
		}
		if err := json.Unmarshal(answers, &raw); err != nil {
			result[order] = &test.StringList{}
		} else {
			parsed := &test.StringList{}
			for _, v := range raw {
				parsed.Values = append(parsed.Values, v.Values...)
			}
			result[order] = parsed
		}
	}

	return &test.GetReadingTrueAnswersBatchResponse{
		TrueAnswers: result,
	}, nil
}

func (t *tests) GetReadingRealTime(ctx context.Context, req *test.ReadingAttemptsRequest) (*test.ReadingAttemptsResponse, error) {

	query := `SELECT end_time FROM r_attempts WHERE user_id = $1 AND id = $2`

	var realEnd time.Time
	err := t.db.QueryRow(query, req.UserId, req.ReadingAtId).Scan(&realEnd)
	if err != nil {
		if req.ReadingAtId <= int32(0) || req.UserId <= int32(0) {
			fmt.Println("attempt id", req.ReadingAtId, "user id", req.UserId)
			return nil, fmt.Errorf("invalid attempt id or user id")
		} else if err == sql.ErrNoRows {
			return nil, fmt.Errorf("real time not found")
		}
		return nil, err
	}

	fmt.Println("real end ===>", realEnd)

	return &test.ReadingAttemptsResponse{
		End: timestamppb.New(realEnd),
	}, nil
}

func (t *tests) SaveReadingResult(ctx context.Context, req *test.ReadingResultRequest) (*test.ReadingResponse, error) {
	if req == nil || len(req.Results) == 0 {
		return nil, errors.New("no results")
	}

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := "INSERT INTO r_results (attempt_id, user_answer, score, is_correct, user_id, test_id) VALUES "
	values := []interface{}{}
	placeholders := []string{}

	for i, result := range req.Results {

		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6),
		)

		values = append(values,
			req.AttemptId,
			result.Answer,
			req.Score,
			result.IsCorrect,
			req.UserId,
			req.TestId,
		)
	}

	if len(placeholders) == 0 {
		return nil, errors.New("no values to insert")
	}

	query += strings.Join(placeholders, ",") + ";"

	_, err = tx.Exec(query, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ReadingResponse{
		Message: "success",
	}, nil

}

func (t *tests) SaveReadingAttempts(ctx context.Context, req *test.ReadingAttemptsRequest) (*test.ReadingAttemptsResponse, error) {

	var attemptID int

	ids := make(pq.Int64Array, len(req.SelectedPessage))
	for i, id := range req.SelectedPessage {
		ids[i] = int64(id)
	}

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	err = tx.QueryRow(`
    INSERT INTO r_attempts (user_id, start_time, end_time)
    VALUES ($1, $2, $3)
    RETURNING id`,
		req.UserId, req.Start.AsTime(), req.End.AsTime()).Scan(&attemptID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ReadingAttemptsResponse{
		Message: "success",
	}, nil
}

func (t *tests) RGetLastAttemptId(ctx context.Context, req *test.RGetLastestAttemptRequest) (*test.RGetLastestAttemptResponse, error) {
	var attemptCount int

	query := `SELECT attempt_id FROM r_results WHERE user_id = $1 order by attempt_id DESC LIMIT 1`

	err := t.db.QueryRow(query, req.UserId).Scan(&attemptCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &test.RGetLastestAttemptResponse{
		LastAttemptId: int32(attemptCount),
	}, nil

}

func (t *tests) CreateListeningQuestions(ctx context.Context, req *test.ListeningQuestionRequest) (*test.ListeningResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	jsonAnswers, err := json.Marshal(req.TrueAnswers)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO l_questions(section_id, order_id, true_answers, content_id) VALUES($1, $2, $3, $4, $5)`

	_, err = tx.Exec(
		query,
		req.SectionId,
		req.OrderId,
		string(jsonAnswers),
		req.ContentId,
		req.TestId,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ListeningResponse{
		Message: "success",
	}, nil
}

func (t *tests) CreateListeningContent(ctx context.Context, req *test.ListeningContentRequest) (*test.ListeningResponse, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO l_contents(section_id, test_id, body) VALUES($1, $2, $3)`

	_, err = tx.Exec(
		query,
		req.SectionId,
		req.TestId,
		req.Body,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ListeningResponse{
		Message: "success",
	}, nil
}

func (t *tests) GetListeningContent(ctx context.Context, req *test.GetListeningContentRequest) (*test.GetListeningContentResponse, error) {

	SBodies := make(map[string]string)

	for _, id := range req.Ids {
		var body string
		err := t.db.QueryRow(`SELECT body FROM l_contents WHERE test_id = $1 AND section_id = $2`,
			req.TestId, id).Scan(&body)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			} else if id > 4 {
				return nil, nil
			}
			return nil, err
		}
		SBodies[fmt.Sprintf("section%d", id)] = body
	}

	return &test.GetListeningContentResponse{
		Bodies: SBodies,
	}, nil

}

func (t *tests) GetListeningTrueAnswers(ctx context.Context, req *test.GetListeningTrueAnswersBatchRequest) (*test.GetListeningTrueAnswersBatchResponse, error) {

	args := make([]interface{}, len(req.OrderIds))
	for i, v := range req.OrderIds {
		args[i] = v

	}

	query := `SELECT order_id, true_answers
		FROM l_questions
		WHERE order_id = ANY($1) AND
		section_id = ANY($2) AND WHERE test_id = $3`

	rows, err := t.db.Query(query, pq.Array(req.OrderIds), pq.Array(req.SectionIds), req.TestId)
	if err != nil {
		fmt.Println("HIHIHIHA --->", err)
		return nil, err
	}

	defer rows.Close()

	result := make(map[int32]*test.StringList)

	for rows.Next() {
		var order int32
		var answers []byte

		if err := rows.Scan(&order, &answers); err != nil {
			fmt.Println("HIHIHIHA ->", err)
			return nil, err
		}
		var raw map[string]struct {
			Values []string `json:"values"`
		}

		if err := json.Unmarshal(answers, &raw); err != nil {
			result[order] = &test.StringList{}
		}
		parsed := &test.StringList{}
		for _, v := range raw {
			parsed.Values = append(parsed.Values, v.Values...)
		}
		result[order] = parsed

	}

	return &test.GetListeningTrueAnswersBatchResponse{
		TrueAnswers: result,
	}, nil
}

func (t *tests) GetListeningRealTime(ctx context.Context, req *test.ListeningAttemptsRequest) (*test.ListeningAttemptsResponse, error) {

	query := `SELECT end_time FROM r_attempts WHERE user_id = $1`

	var realEnd *timestamppb.Timestamp
	err := t.db.QueryRow(query, req.UserId).Scan(&realEnd)
	if err != nil {
		return nil, err
	}

	return &test.ListeningAttemptsResponse{
		End: realEnd,
	}, nil
}

func (t *tests) SaveListeningAnswers(ctx context.Context, req *test.ListeningResultRequest) (*test.ListeningResponse, error) {
	if len(req.Results) == 0 {
		return nil, errors.New("no results")
	}

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	query := `INSERT INTO l_results(attempt_id, user_answer, score, is_correct, user_id, test_id) VALUES `
	values := []interface{}{}
	placeholders := []string{}

	for i, results := range req.Results {

		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
		values = append(values,
			req.ListeningAtId,
			results.Answer,
			req.Score,
			results.IsCorrect,
			req.UserId,
			results.TestId,
		)
	}

	if len(placeholders) == 0 {
		return nil, errors.New("no results")
	}

	query += strings.Join(placeholders, ",") + ";"

	_, err = tx.Exec(query, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ListeningResponse{
		Message: "success",
	}, nil
}

func (t *tests) SaveListeningAttempts(ctx context.Context, req *test.ListeningAttemptsRequest) (*test.ListeningResponse, error) {
	var attemptID int

	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	err = tx.QueryRow(`INSERT into l_attempts (user_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING id`, req.UserId, req.Start.AsTime(), req.End.AsTime()).Scan(&attemptID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &test.ListeningResponse{
		Message:   "success",
		AttemptId: int32(attemptID),
	}, nil
}

func (t *tests) LGetLastestAttempt(ctx context.Context, req *test.LLastestAttemptRequest) (*test.LLastestAttemptResponse, error) {
	var attemptCount int

	query := `SELECT attempt_id FROM l_results WHERE user_id = $1 ORDER BY attempt_id DESC LIMIT 1`

	err := t.db.QueryRow(query, req.UserId).Scan(&attemptCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &test.LLastestAttemptResponse{
		LastAttemptId: int32(attemptCount),
	}, nil
}
