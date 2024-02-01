package coursedto

import (
	"github.com/google/uuid"
	coursemodel "server/modules/course/model"
)

type CourseEntity struct {
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	BackgroundImg string             `json:"background_img"`
	Thumbnail     string             `json:"thumbnail"`
	StartDate     *string            `json:"start_date"`
	EndDate       *string            `json:"end_date"`
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
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Lectures    []*SimpleLecture `json:"lectures"`
}

// CREATE COURSE REQUEST

type CourseCreateRequest struct {
	CourseEntity `json:",inline"`
	CourseInfos  []CourseInfoEntity `json:"course_infos"`
	Sections     []SectionEntity    `json:"sections"`
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
		var lectures []*coursemodel.Lecture
		for _, lecture := range section.Lectures {
			lectures = append(lectures, &coursemodel.Lecture{
				Name:        lecture.Name,
				Description: lecture.Description,
			})
		}
		sections = append(sections, &coursemodel.Section{
			Name:        section.Name,
			Description: section.Description,
			Lectures:    lectures,
		})
	}

	return &coursemodel.Course{
		Name:          course.Name,
		Description:   course.Description,
		BackgroundImg: course.BackgroundImg,
		Thumbnail:     course.Thumbnail,
		StartDate:     course.StartDate,
		EndDate:       course.EndDate,
		Price:         course.Price,
		Currency:      course.Currency,
		Level:         coursemodel.CourseLevel(course.Level),
		SubjectId:     course.SubjectId,
		Grade:         course.Grade,
		CourseInfos:   courseInfos,
		Sections:      sections,
	}
}

// ATTEMPT COURSE BY CODE

type CourseAttemptRequest struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type CourseEnrollUpdateRequest struct {
	Status         string `json:"status"`
	CourseEnrollId string `json:"id"`
}
