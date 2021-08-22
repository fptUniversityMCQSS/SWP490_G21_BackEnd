package response

import "time"

type HistoryResponse struct {
	Id                int64     `orm:"pk;auto" json:"id" form:"id"`
	Date              time.Time `json:"historyDate" form:"historyDate"`
	Name              string    `json:"historyName" form:"historyName"`
	User              string    `json:"user"`
	Subject           string    `json:"subject"`
	NumberOfQuestions int64     `json:"questions_number"`
	Status            string    `json:"status"`
}
