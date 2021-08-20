package unity

import "time"

type ExamTest struct {
	Id                int64 `orm:"pk;auto"`
	User              *User `orm:"rel(fk);on_delete(cascade)"`
	Name              string
	Path              string
	Date              time.Time   `json:"test_date" form:"test_date"`
	Questions         []*Question `orm:"reverse(many)"`
	Subject           string
	NumberOfQuestions int64
	Status            string
}
