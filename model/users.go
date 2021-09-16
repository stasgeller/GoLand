package model

import "encoding/xml"

type Users struct {
	Users []User `xml:"user" json:"users"`
}

type User struct {
	XMLName *xml.Name `xml:"user" json:",omitempty"`
	Name    string   `xml:"name" json:"name"`
	Type    string   `xml:"type,attr" json:"type"`
	Social  Social   `xml:"social" json:"social"`
}

type Social struct {
	XMLName  *xml.Name `xml:"social" json:",omitempty"`
	Facebook string   `xml:"facebook" json:"facebook"`
	Twitter  string   `xml:"twitter" json:"twitter"`
}
