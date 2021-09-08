package parsers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Users struct {
	Users []User `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
}

func ParseXml(w http.ResponseWriter, r *http.Request) {
	xmlFile, err := os.Open("./parsers/users.xml")

	if err != nil {
		fmt.Fprintln(w, "ERROR: "+err.Error())
		return
	}

	defer xmlFile.Close()
	byteFile, _ := ioutil.ReadAll(xmlFile)
	var users Users

	xml.Unmarshal(byteFile, &users)

	json.NewEncoder(w).Encode(users)
}
