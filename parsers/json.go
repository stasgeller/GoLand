package parsers

import (
	"api.com/model"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
)

type JsonParser struct{}

func (j JsonParser) Parse(file multipart.File, header multipart.FileHeader) (*model.Users, error) {
	var users model.Users
	byteFile, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteFile, &users)

	return &users, nil
}

func NewJsonParser() Parser {
	return JsonParser{}
}
