package response

import "time"

type KnowledgResponse struct {
	Id       int64     `json:"knowledgeId"`
	Name     string    `json:"knowledgeName" form:"knowledgeName"`
	Date     time.Time `orm:"auto_now_add" json:"knowledgeDate" form:"knowledgeDate"`
	Username string    `json:"Username" form:"Username"`
	Status   string    `json:"status" form:"status"`
}
