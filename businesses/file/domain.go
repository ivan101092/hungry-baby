package file

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	Type       string
	URL        string
	UserUpload string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, fileDomain *Domain) (Domain, error)
	Update(ctx context.Context, fileDomain *Domain) (Domain, error)
	Delete(ctx context.Context, fileDomain *Domain) (Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	FindByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, fileDomain *Domain) (Domain, error)
	Update(ctx context.Context, fileDomain *Domain) (Domain, error)
	Delete(ctx context.Context, fileDomain *Domain) (Domain, error)
}
