package assignmentrecognizeprovider

import "context"

type Provider interface {
	RecognizeAssignment(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error)
}
