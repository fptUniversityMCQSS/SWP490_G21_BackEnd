package model

type Question struct {
	Id       int64 `orm:"pk;auto"`
	Number   int64
	Mark     float64
	Content  string    `json:"question_content" form:"question_content"`
	Options  []*Option `orm:"reverse(many)"`
	Answer   *Option   `orm:"null;rel(one);on_delete(set_null)" json:"answer" form:"answer"`
	ExamTest *ExamTest `orm:"rel(fk)"`
}
