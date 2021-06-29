package response

type QuestionResponse struct {
	Number  int64
	Content string
	Options []OptionResponse
	Answer  OptionResponse
}
