package response

type QuestionResponse struct {
	Number  int64
	Content string
	Options []OptionResponse
	Answer  string
}

type QuestionAnswerResponse struct {
	Qn     int64  `json:"qn"`
	Answer string `json:"answer"`
}
