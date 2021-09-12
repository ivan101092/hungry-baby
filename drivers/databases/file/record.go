package file

import (
	"hungry-baby/businesses/file"
	fileUsecase "hungry-baby/businesses/file"
	"hungry-baby/drivers/postgres"
)

type File struct {
	ID         int
	Type       string
	URL        string
	UserUpload string
	postgres.BaseModel
}

func FromDomain(domain *fileUsecase.Domain) *File {
	return &File{
		ID:         domain.ID,
		Type:       domain.Type,
		URL:        domain.URL,
		UserUpload: domain.UserUpload,
	}
}

func (rec *File) ToDomain() file.Domain {
	return file.Domain{
		ID:         rec.ID,
		Type:       rec.Type,
		URL:        rec.URL,
		UserUpload: rec.UserUpload,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}
