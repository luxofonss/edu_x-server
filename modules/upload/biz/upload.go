package uploadbiz

import (
	"context"
	"fmt"
	"image"
	"io"
	"path/filepath"
	"strings"
	"time"

	"server/common"
	uploadprovider "server/libs/upload_provider"
	uploadmodel "server/modules/upload/model"
)

type CreateImageRepo interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.Provider
	imgRepo  CreateImageRepo
}

func NewUploadBiz(provider uploadprovider.Provider, imgRepo CreateImageRepo) *uploadBiz {
	return &uploadBiz{provider: provider, imgRepo: imgRepo}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	//fileBytes := bytes.NewBuffer(data)
	//
	//w, h, err := getImageDimension(fileBytes)
	//if err != nil {
	//	return nil, uploadmodel.ErrFileIsNotImage(err)
	//}

	if strings.TrimSpace(folder) == "" {
		folder = "files"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUpload(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}
	//
	//img.Width = w
	//img.Height = h
	img.Extension = fileExt
	img.CloudName = biz.provider.GetCloudName()

	// save to db

	return img, nil

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
