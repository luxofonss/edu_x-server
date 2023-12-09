package uploadprovider

import (
	"context"

	"server/common"
)

type Provider interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error)
	GetCloudName() string
}
