package entity

type Option struct {
	Id         int64     `orm:"pk;auto"`
	QuestionId *Question `orm:"rel(fk);on_delete(cascade)"`
	Key        string    `json:"OptionKey" form:"OptionKey"`
	Content    string    `json:"OptionContent" form:"OptionContent"`
}
