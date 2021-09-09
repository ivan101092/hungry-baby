package file

import (
	"context"
	"mime/multipart"
	"time"
)

type Domain struct {
	ID         int
	Type       string
	URL        string
	FullURL    string
	UserUpload string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, fileDomain *Domain) (Domain, error)
	Upload(ctx context.Context, types, filePath string, f *multipart.FileHeader) (Domain, error)
	Delete(ctx context.Context, fileDomain *Domain) (Domain, error)
}

type Repository interface {
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, fileDomain *Domain) (Domain, error)
	Delete(ctx context.Context, fileDomain *Domain) (Domain, error)
}
