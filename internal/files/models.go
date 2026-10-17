package files

import (
	"time"

	"github.com/twonumberfortyfives/home-cloud-backend/internal/users"
)

type FileDto struct {
	Name      string
	Size      int64
	Icon      string
	CreatedBy users.UserDto
	UpdatedAt time.Time
	CreatedAt time.Time
}

type FileActionRequest struct {
	Path string
}

type FileListResponse struct {
	Files []FileDto
}
