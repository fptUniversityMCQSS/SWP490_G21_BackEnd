package response

import "time"

type ExamTestResponse struct {
	Id                int64
	Name              string
	User              UserResponse
	Subject           string
	NumberOfQuestions int64
	Date              time.Time
	Questions         []QuestionResponse
}
