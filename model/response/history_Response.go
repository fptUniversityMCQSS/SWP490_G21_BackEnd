package response

import "time"

type HistoryResponse struct {
	Id   int64     `orm:"pk;auto" json:"id" form:"id"`
	Date time.Time `orm:"auto_now_add" json:"historyDate" form:"historyDate"`
	Name string    `json:"historyName" form:"historyName"`
}
