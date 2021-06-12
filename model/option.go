package model

type Option struct {
	Id *int64      `orm:"pk"`
	QuestionId *Question  `orm:"rel(fk)"`
	Key       string
	Content   string
	Paragraph string
}
