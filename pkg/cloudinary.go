package pkg

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func Cloudinary(file interface{}) (string, error) {
	name := viper.GetString("cloudinary.cloud")
	key := viper.GetString("cloudinary.key")
	secret := viper.GetString("cloudinary.secret")

	cld, _ := cloudinary.NewFromParams(name, key, secret)

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
