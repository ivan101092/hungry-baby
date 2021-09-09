package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Connection ..
type Connection struct {
	AccessKey string
	SecretKey string
	UseSSL    bool
	BaseURL   string
	Duration  int
	Bucket    string
}

// InitClient ...
func (conn *Connection) InitClient() (client *minio.Client, err error) {
	client, err = minio.New(conn.BaseURL, &minio.Options{
		Creds:  credentials.NewStaticV4(conn.AccessKey, conn.SecretKey, ""),
		Secure: conn.UseSSL,
	})
	if err != nil {
		return client, err
	}

	return client, nil
}
