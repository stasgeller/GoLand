package parsers

import (
	"api.com/model"
	"encoding/xml"
	"io/ioutil"
	"mime/multipart"
)

type XmlParser struct {
}

func (x XmlParser) Parse(file multipart.File) (*model.Users, error) {
	var users model.Users
	byteFile, _ := ioutil.ReadAll(file)

	xml.Unmarshal(byteFile, &users)

	return &users, nil
}

func NewXmlParser() Parser {
	return XmlParser{}
}
