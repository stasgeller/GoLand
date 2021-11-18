package parsers

import (
	"mime/multipart"

	"api.com/model"
)

type Parser interface {
	Parse(file multipart.File) (*model.Users, error)
}
