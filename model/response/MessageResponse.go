package response

type Message struct {
	Message string `json:"message"`
}

func (e Message) Error() string {
	return e.Message
}
