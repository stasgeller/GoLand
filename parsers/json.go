package parsers

import (
	"api.com/model"
	"encoding/json"
	"mime/multipart"
)

type JsonParser struct{}

func NewJsonParser() Parser {
	return JsonParser{}
}

func (j JsonParser) Parse(file multipart.File) (*model.Users, error) {
	var users model.Users
	err := json.NewDecoder(file).Decode(&users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}
