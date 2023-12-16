package assignmentmodel

import (
	"github.com/google/uuid"
	"server/common"
)

const QuestionEntityName = "Question"

type QuestionType string

const (
	MultipleChoice QuestionType = "multiple_choice"
	SingleChoice   QuestionType = "single_choice"
	ShortAnswer    QuestionType = "short_answer"
	LongAnswer     QuestionType = "long_answer"
	SuperQuestion  QuestionType = "super_question"
)

type QuestionLevel string

const (
	Easy      QuestionLevel = "easy"
	Medium    QuestionLevel = "medium"
	Hard      QuestionLevel = "hard"
	SuperHard QuestionLevel = "super_hard"
)

type Question struct {
	common.SQLModel `json:",inline"`
	Title           string                   `json:"title" gorm:"column:title;"`
	Image           *common.Image            `json:"image" gorm:"column:image;"`
	AudioUrl        string                   `json:"audio_url" gorm:"column:audio_url;"`
	Type            QuestionType             `json:"type" gorm:"column:type;"`
	Level           QuestionLevel            `json:"level" gorm:"column:level;"`
	AssignmentId    uuid.UUID                `json:"assignment_id" gorm:"column:assignment_id;type:uuid;"`
	TeacherId       uuid.UUID                `json:"teacher_id" gorm:"column:teacher_id;type:uuid;default:NULL;"`
	SchoolId        uuid.UUID                `json:"school_id" gorm:"column:school_id;type:uuid;default:NULL;"`
	SubjectId       uuid.UUID                `json:"subject_id" gorm:"column:subject_id;type:uuid;"`
	Order           *int                     `json:"order" gorm:"-"`
	Point           *int                     `json:"point" gorm:"-"`
	ParentId        *uuid.UUID               `json:"parent_id" gorm:"column:parent_id;type:uuid;"`
	Choices         []*QuestionChoice        `json:"choices" gorm:"foreignKey:QuestionId"`
	CorrectAnswers  []*QuestionCorrectAnswer `json:"correct_answers" gorm:"foreignKey:QuestionId;"`
	Answers         []*QuestionAnswer        `json:"answer" gorm:"foreignKey:QuestionId"`
	Questions       []*Question              `json:"questions" gorm:"foreignKey:ParentId"`
}

func (Question) TableName() string { return "questions" }
