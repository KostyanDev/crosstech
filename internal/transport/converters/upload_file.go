package converters

import "app/internal/domain"

type UploadFileRequest struct {
	FileName string `json:"file_name"`
}

func ToDomainUploadFileName(file UploadFileRequest) domain.File {
	return domain.File{
		Name: file.FileName,
	}
}
