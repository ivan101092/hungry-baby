package minio

import (
	"context"
	minioBusiness "hungry-baby/businesses/minio"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/minio/minio-go/v7"
)

type MinioModel struct {
	Client *minio.Client
	Bucket string
}

const defaultDuration = 15

// NewMinioModel ...
func NewMinioModel(client *minio.Client, bucket string) minioBusiness.Repository {
	return &MinioModel{
		Client: client,
		Bucket: bucket,
	}
}

// Upload ...
func (model *MinioModel) Upload(path string, fileHeader *multipart.FileHeader) (res string, err error) {
	src, err := fileHeader.Open()

	if err != nil {
		return res, err
	}
	defer src.Close()

	fileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")
	path += `/local/` + fileName

	_, err = model.Client.PutObject(context.Background(), model.Bucket, path, src, fileHeader.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return res, err
	}
	res = path
	return res, nil
}

// GetFile ...
func (model *MinioModel) GetFile(objectName string) (res string, err error) {
	reqParams := make(url.Values)

	duration := time.Minute * defaultDuration
	uri, err := model.Client.PresignedGetObject(context.Background(), model.Bucket, objectName, duration, reqParams)
	if err != nil {
		return res, err
	}
	res = uri.String()

	return res, err
}

// Delete ...
func (model *MinioModel) Delete(objectName string) (err error) {
	options := minio.RemoveObjectOptions{}
	err = model.Client.RemoveObject(context.Background(), model.Bucket, objectName, options)

	return err
}
