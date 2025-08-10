package wb2pages

import (
	"compress/bzip2"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

const Prefix string = `<root
	xmlns="http://www.mediawiki.org/xml/export-0.11/"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.mediawiki.org/xml/export-0.11/ http://www.mediawiki.org/xml/export-0.11.xsd"
	version="0.11"
	xml:lang="en"
>`

var PrefixBytes []byte = []byte(Prefix)
var SuffixBytes []byte = []byte("</root>")

type BasicPage struct {
	XMLName  xml.Name `xml:"page"`
	Title    string   `xml:"title,omitempty"`
	Ns       string   `xml:"ns,omitempty"`
	Id       string   `xml:"id,omitempty"`
	Redirect Redirect `xml:"redirect,omitempty"`
	Revision Revision `xml:"revision,omitempty"`
}

func (b BasicPage) ShortString() string {
	return fmt.Sprintf("Page{Title:%q, Ns:%s, Id:%s}", b.Title, b.Ns, b.Id)
}

type Redirect struct {
	Title string `xml:"title,attr,omitempty"`
}

type Revision struct {
	Id          string      `xml:"id,omitempty"`
	ParentId    string      `xml:"parentid,omitempty"`
	Timestamp   string      `xml:"timestamp,omitempty"`
	Contributor Contributor `xml:"contributor,omitempty"`
	Comment     string      `xml:"comment,omitempty"`
	Origin      string      `xml:"origin,omitempty"`
	Model       string      `xml:"model,omitempty"`
	Format      string      `xml:"format,omitempty"`
	Text        Text        `xml:"text,omitempty"`
}

func (r Revision) ShortString() string {
	return fmt.Sprintf(
		"Revision{Id:%s, Timestamp:%s, Contributor:%s}",
		r.Id,
		r.Timestamp,
		r.Contributor.Username,
	)
}

type Contributor struct {
	Username string `xml:"username,omitempty"`
	Id       string `xml:"id,omitempty"`
}

type Text struct {
	XMLName  xml.Name `xml:"text"`
	Bytes    string   `xml:"bytes,attr,omitempty"`
	Sha1     string   `xml:"sha1,attr,omitempty"`
	XMLSpace string   `xml:"xml:space,attr,omitempty"`
	Value    string   `xml:",chardata"`
}

func BasicPages(pagesText string) ([]BasicPage, error) {
	var fullXML string = Prefix + pagesText + "</root>"

	var root struct {
		Pages []BasicPage `xml:"page"`
	}

	if err := xml.Unmarshal([]byte(fullXML), &root); err != nil {
		return nil, err
	}

	return root.Pages, nil
}

func BasicPagesBytes(pagesText []byte) ([]BasicPage, error) {
	var fullXml []byte
	fullXml = append(fullXml, PrefixBytes...)
	fullXml = append(fullXml, pagesText...)
	fullXml = append(fullXml, SuffixBytes...)

	var root struct {
		Pages []BasicPage `xml:"page"`
	}

	if err := xml.Unmarshal(fullXml, &root); err != nil {
		return nil, err
	}

	return root.Pages, nil
}

// Converts the byte range to pages.
// TODO: reduce allocations.
func ReaderToPages(bzip2rdr io.ReaderAt, offset int64, size int64) ([]BasicPage, error) {
	section := io.NewSectionReader(bzip2rdr, offset, size)
	bz2Reader := bzip2.NewReader(section)
	decompressed, err := io.ReadAll(bz2Reader)
	if err != nil {
		return nil, err
	}
	return BasicPagesBytes(decompressed)
}

// Converts the byte range of the file to pages.
// TODO: reduce allocations.
func FileToPages(f *os.File, offset int64, size int64) ([]BasicPage, error) {
	return ReaderToPages(f, offset, size)
}
