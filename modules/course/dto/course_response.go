package coursedto

import (
	"fmt"
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

func ToSimpleSubject(subject coursemodel.Subject) *SimpleSubject {
	return &SimpleSubject{
		Id:   subject.Id.String(),
		Name: subject.Name,
	}
}

type SimpleLecture struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToSimpleLecture(lecture coursemodel.Lecture) SimpleLecture {
	return SimpleLecture{
		Id:          lecture.Id.String(),
		Name:        lecture.Name,
		Description: lecture.Description,
	}
}

type SimpleSection struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Lectures    []SimpleLecture `json:"lectures"`
}

func ToSimpleSection(section coursemodel.Section) *SimpleSection {
	var lectures []SimpleLecture
	for _, lecture := range section.Lectures {
		lectures = append(lectures, ToSimpleLecture(*lecture))
	}
	return &SimpleSection{
		Id:          section.Id.String(),
		Name:        section.Name,
		Description: section.Description,
		Lectures:    lectures,
	}
}

type CourseInfo struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func ToCourseInfo(courseInfo coursemodel.CourseInfo) CourseInfo {
	return CourseInfo{
		Content: courseInfo.Content,
		Type:    string(courseInfo.Type),
	}
}

type SimpleCourseResponse struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	BackgroundImg string         `json:"background_img"`
	Thumbnail     string         `json:"thumbnail"`
	Price         float64        `json:"price"`
	Level         string         `json:"level"`
	Currency      string         `json:"currency"`
	Grade         int            `json:"grade"`
	Code          string         `json:"code"`
	Teacher       *SimpleTeacher `json:"teacher"`
	Subject       *SimpleSubject `json:"subject"`
}

type DetailCourseResponse struct {
	Id            string                    `json:"id"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	BackgroundImg string                    `json:"background_img"`
	Thumbnail     string                    `json:"thumbnail"`
	Price         float64                   `json:"price"`
	Level         string                    `json:"level"`
	Currency      string                    `json:"currency"`
	Grade         int                       `json:"grade"`
	Code          string                    `json:"code"`
	Teacher       *SimpleTeacher            `json:"teacher"`
	Subject       *SimpleSubject            `json:"subject"`
	Sections      []*SimpleSection          `json:"sections"`
	CourseInfos   []*coursemodel.CourseInfo `json:"course_infos"`
}

func ToSimpleCourseResponse(course coursemodel.Course) SimpleCourseResponse {
	return SimpleCourseResponse{
		Id:            course.Id.String(),
		Name:          course.Name,
		Description:   course.Description,
		BackgroundImg: course.BackgroundImg,
		Thumbnail:     course.Thumbnail,
		Price:         course.Price,
		Level:         string(course.Level),
		Currency:      course.Currency,
		Grade:         course.Grade,
		Code:          course.Code,
		Teacher:       ToSimpleTeacher(*course.Teacher),
		Subject:       ToSimpleSubject(*course.Subject),
	}
}

func ToDetailCourseResponse(course coursemodel.Course) DetailCourseResponse {
	var sections []*SimpleSection
	for _, section := range course.Sections {
		sections = append(sections, ToSimpleSection(*section))
	}

	fmt.Println("course.CourseInfos:", course.CourseInfos)

	var courseInfos []*coursemodel.CourseInfo
	for _, courseInfo := range course.CourseInfos {
		courseInfos = append(courseInfos, courseInfo)
	}

	return DetailCourseResponse{
		Id:            course.Id.String(),
		Name:          course.Name,
		Description:   course.Description,
		BackgroundImg: course.BackgroundImg,
		Thumbnail:     course.Thumbnail,
		Price:         course.Price,
		Level:         string(course.Level),
		Currency:      course.Currency,
		Grade:         course.Grade,
		Code:          course.Code,
		Teacher:       ToSimpleTeacher(*course.Teacher),
		Subject:       ToSimpleSubject(*course.Subject),
		Sections:      sections,
		CourseInfos:   courseInfos,
	}
}
