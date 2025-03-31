package atom

import "encoding/xml"

type AtomFeed struct {
	XMLName xml.Name `xml:"feed"`
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated string   `xml:"updated"`
	Link    []Link   `xml:"link"`
	Entries []Entry  `xml:"entry"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

type Entry struct {
	ID        string `xml:"id"`
	Title     string `xml:"title"`
	Updated   string `xml:"updated"`
	Published string `xml:"published"`
	Link      Link   `xml:"link"`
	Content   string `xml:"content"`
	Author    Author `xml:"author"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}
