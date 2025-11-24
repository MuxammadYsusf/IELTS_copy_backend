package models

import "time"

type Questions struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	A         string `json:"a"`
	B         string `json:"b"`
	TruAnsver string `json:"true_answer"`
	SubjectId int32  `json:"subject_id"`
	GradeId   int32  `json:"grade_id"`
}

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phonenumber"`
	Password    string `json:"password"`
}

type Result struct {
	ResultId       int    `json:"id"`
	UserId         int    `json:"user_id"`
	QuestionId     int    `json:"question_id"`
	AttemptCount   int    `json:"attempt_id"`
	IsCorrect      bool   `json:"is_correct"`
	SelectedOption string `json:"selected_option"`
}

type Attempts struct {
	UserId         int    `json:"user_id"`
	A              string `json:"a"`
	B              string `json:"b"`
	AttemptId      int    `json:"attempt_id"`
	QuestionId     int    `json:"question_id"`
	SelectedOption string `json:"selected_option"`
	IsCorrect      bool   `json:"is_correct"`
	SubjectId      int32  `json:"subject_id"`
	GradeId        int32  `json:"grade_id"`
}

type AttemptList struct {
	Id           int       `json:"attempt_id"`
	UserId       int       `json:"user_id"`
	Created_at   time.Time `json:"created_at"`
	CorrectCount int       `json:"correct_answers"`
}
