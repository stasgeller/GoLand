package parsers

import (
	"api.com/model"
	"encoding/json"
	"mime/multipart"
)

type XmlParser struct {
}

func NewXmlParser() Parser {
	return XmlParser{}
}

func (x XmlParser) Parse(file multipart.File) (*model.Users, error) {
	var users model.Users
	err := json.NewDecoder(file).Decode(&users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}
