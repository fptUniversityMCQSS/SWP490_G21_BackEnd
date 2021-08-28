package entity

type Question struct {
	Id       int64 `orm:"pk;auto"`
	Number   int64
	Mark     float64
	Content  string    `json:"Content" form:"Content"`
	Options  []*Option `orm:"reverse(many)"`
	Answer   string
	ExamTest *ExamTest `orm:"rel(fk);on_delete(cascade)"`
}
