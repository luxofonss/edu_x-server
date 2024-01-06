package assignmentbiz

import (
	"context"
	assignmentrecognizeprovider "server/libs/assignment_recognize_provider"
)

type recognizeAssignmentBiz struct {
	provider assignmentrecognizeprovider.Provider
}

func NewRecognizeAssignmentBiz(provider assignmentrecognizeprovider.Provider) *recognizeAssignmentBiz {
	return &recognizeAssignmentBiz{provider: provider}
}

func (biz *recognizeAssignmentBiz) RecognizeAssignment(context context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	return biz.provider.RecognizeAssignment(context, data)
}
