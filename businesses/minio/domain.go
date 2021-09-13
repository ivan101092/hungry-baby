package minio

import "mime/multipart"

type Domain struct {
}

type Repository interface {
	Upload(path string, file *multipart.FileHeader) (string, error)
	GetFile(objectName string) (string, error)
	Delete(objectName string) error
}
