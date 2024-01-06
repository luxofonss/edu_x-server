package coursedto

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

type CourseEntity struct {
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	BackgroundImg string             `json:"background_img"`
	StartDate     string             `json:"start_date"`
	EndDate       string             `json:"end_date"`
	Price         float64            `json:"price"`
	Currency      string             `json:"currency"`
	Level         string             `json:"level"`
	SubjectId     uuid.UUID          `json:"subject_id"`
	Grade         int                `json:"grade"`
	CourseInfos   []CourseInfoEntity `json:"course_infos"`
}

type CourseInfoEntity struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type SectionEntity struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CREATE COURSE REQUEST

type CourseCreateRequest struct {
	Course      CourseEntity       `json:"course"`
	CourseInfos []CourseInfoEntity `json:"infos"`
	Sections    []SectionEntity    `json:"sections"`
}

func (course CourseCreateRequest) ToCourseModel() *coursemodel.Course {
	var courseInfos []*coursemodel.CourseInfo
	var sections []*coursemodel.Section

	for _, info := range course.CourseInfos {
		courseInfos = append(courseInfos, &coursemodel.CourseInfo{
			Content: info.Content,
			Type:    coursemodel.CourseInfoTypes(info.Type),
		})
	}

	for _, section := range course.Sections {
		sections = append(sections, &coursemodel.Section{
			Name:        section.Name,
			Description: section.Description,
		})
	}

	return &coursemodel.Course{
		Name:          course.Course.Name,
		Description:   course.Course.Description,
		BackgroundImg: course.Course.BackgroundImg,
		StartDate:     course.Course.StartDate,
		EndDate:       course.Course.EndDate,
		Price:         course.Course.Price,
		Currency:      course.Course.Currency,
		Level:         coursemodel.CourseLevel(course.Course.Level),
		SubjectId:     course.Course.SubjectId,
		Grade:         course.Course.Grade,
		CourseInfos:   courseInfos,
		Sections:      sections,
	}
}

// ATTEMPT COURSE BY CODE

type CourseAttemptRequest struct {
	Code string `json:"code" form:"code" binding:"required"`
}
