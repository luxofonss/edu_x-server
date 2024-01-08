package coursedto

import (
	"server/common"
	coursemodel "server/modules/course/model"
)

type SimpleTeacher struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func ToSimpleTeacher(teacher common.SimpleUser) *SimpleTeacher {
	return &SimpleTeacher{
		Id:        teacher.Id.String(),
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
	}
}

type SimpleSubject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ToSimpleSubject(subject coursemodel.SimpleSubjectGet) *SimpleSubject {
	return &SimpleSubject{
		Id:   subject.Id.String(),
		Name: subject.Name,
	}
}

type SimpleCourseResponse struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	BackgroundImg string         `json:"background_img"`
	Price         float64        `json:"price"`
	Level         string         `json:"level"`
	Currency      string         `json:"currency"`
	Grade         int            `json:"grade"`
	Teacher       *SimpleTeacher `json:"teacher"`
	Subject       *SimpleSubject `json:"subject"`
}

func ToSimpleCourseResponse(course coursemodel.CourseGet) SimpleCourseResponse {
	return SimpleCourseResponse{
		Id:            course.Id.String(),
		Name:          course.Name,
		Description:   course.Description,
		BackgroundImg: course.BackgroundImg,
		Price:         course.Price,
		Level:         string(course.Level),
		Currency:      course.Currency,
		Grade:         course.Grade,
		Teacher:       ToSimpleTeacher(*course.Teacher),
		Subject:       ToSimpleSubject(*course.Subject),
	}
}
