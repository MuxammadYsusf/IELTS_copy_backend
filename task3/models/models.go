package models

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phonenumber"`
	Password    string `json:"password"`
}

type ReadingQuestions struct {
	Id         int    `json:"id"`
	Question   string `json:"question"`
	PessageId  int    `json:"pessage"`
	TestId     int    `json:"test"`
	OrderId    int    `json:"order"`
	Body       string `json:"body"`
	TrueAnswer string `json:"true_answer"`
}

type ListeningQuestions struct {
	Id         int    `json:"id"`
	Question   string `json:"question"`
	SectionId  int    `json:"section"`
	TestId     int    `json:"test"`
	OrderId    int    `json:"order"`
	Body       string `json:"body"`
	TrueAnswer string `json:"true_answer"`
}
