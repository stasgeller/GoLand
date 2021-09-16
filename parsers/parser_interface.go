package parsers

import (
	"mime/multipart"

	"api.com/model"
)

type Parser interface {
	Parse(file multipart.File, header multipart.FileHeader) (*model.Users, error)
}
