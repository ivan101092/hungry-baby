package request

import "hungry-baby/businesses/file"

type File struct {
	Type       string `json:"type"`
	URL        string `json:"url"`
	UserUpload string `json:"user_upload"`
}

func (req *File) ToDomain() *file.Domain {
	return &file.Domain{
		Type:       req.Type,
		URL:        req.URL,
		UserUpload: req.UserUpload,
	}
}
