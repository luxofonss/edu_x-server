package coursedto

type CourseEnrollUpdateRequest struct {
	Status   string `json:"status"`
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
}
