package response

import "time"

type ExamTestResponse struct {
	Id        int64
	Name      string
	User      UserResponse
	Date      time.Time
	Questions []QuestionResponse
}
