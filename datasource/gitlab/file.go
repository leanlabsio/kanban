package gitlab

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
)

// UploadFile uploads file to gitlab
func (ds GitLabDataSource) UploadFile(boardID string, file models.UploadForm) (*models.File, error) {
	body, err := file.File.Open()
	res, err := ds.client.UploadFile(boardID, file.File.Filename, body)

	return newFileFromGitlab(res), err
}

// newFileFromGitlab creates new local file from gitlab file
func newFileFromGitlab(f *gitlab.File) *models.File {
	return &models.File{
		Alt:      f.Alt,
		URL:      f.URL,
		Markdown: f.Markdown,
	}
}
