package response

import "time"

type KnowledgResponse struct {
	Name string    `json:"knowledgeName" form:"knowledgeName"`
	Date time.Time `orm:"auto_now_add" json:"knowledgeDate" form:"knowledgeDate"`
	Username string  `json:"Username" form:"Username"`
}
