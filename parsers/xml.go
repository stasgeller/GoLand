package parsers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"api.com/model"
)

type XmlParser struct {
}

func (x XmlParser) Parse(file multipart.File, header multipart.FileHeader) (*model.Users, error) {
	pointer, err := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(fmt.Sprintf("Read file error: %s", err))
	}

	defer pointer.Close()

	var users model.Users
	byteFile, _ := ioutil.ReadAll(file)

	xml.Unmarshal(byteFile, &users)

	return &users, nil
}

func NewXmlParser() Parser {
	return XmlParser{}
}
