package models

import "mime/multipart"

// File represents uploaded file
type File struct {
	Alt      string `json:"alt"`
	URL      string `json:"url"`
	Markdown string `json:"markdown"`
}

// UploadForm represents file uploaded
type UploadForm struct {
	File *multipart.FileHeader `form:"file"`
}
