package response

import (
	"hungry-baby/businesses/file"
	"time"
)

type File struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	URL        string    `json:"url"`
	FullURL    string    `json:"full_url"`
	UserUpload string    `json:"user_upload"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain file.Domain) File {
	return File{
		ID:         domain.ID,
		Type:       domain.Type,
		URL:        domain.URL,
		FullURL:    domain.FullURL,
		UserUpload: domain.UserUpload,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
