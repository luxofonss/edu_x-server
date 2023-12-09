package assignmentmodel

import "server/common"

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
	TeacherId       *int                     `json:"teacher_id" gorm:"column:teacher_id;"`
	SchoolId        *int                     `json:"school_id" gorm:"column:school_id;"`
	SubjectId       int                      `json:"subject_id" gorm:"column:subject_id;"`
	Order           *int                     `json:"order" gorm:"-"`
	Point           *int                     `json:"point" gorm:"-"`
	Questions       []*QuestionCreate        `json:"questions" gorm:"-"`
	Choices         []*QuestionChoice        `json:"choices" gorm:"foreignKey:QuestionId"`
	CorrectAnswers  []*QuestionCorrectAnswer `json:"correct_answers" gorm:"foreignKey:QuestionId"`
}

func (Question) TableName() string { return "questions" }

func (q *Question) Mask(isAdminOrOwner bool) {
	q.GenUID(common.DbTypeQuestion)
}

type QuestionCreate struct {
	Question      Question                `json:"question"`
	Choices       []QuestionChoice        `json:"choices"`
	CorrectAnswer []QuestionCorrectAnswer `json:"correct_answers"`
}

type CreatedQuestionRelation struct {
	QuestionId int `json:"question_id"`
	Order      int `json:"order"`
	Point      int `json:"point"`
}
