package model

type Option struct {
	Id         int64     `orm:"pk;auto"`
	QuestionId *Question `orm:"rel(fk)"`
	Key        string    `json:"OptionKey" form:"OptionKey"`
	Content    string    `json:"OptionContent" form:"OptionContent"`
	Paragraph  string    `json:"OptionParagraph" form:"OptionParagraph"`
}
