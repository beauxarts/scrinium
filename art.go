package litres_integration

import (
	"encoding/xml"
)

type LitResUpdates struct {
	XMLName xml.Name `xml:"litres-updates"`
	Arts    []Art    `xml:"art"`
}

type Art struct {
	XMLName     xml.Name   `xml:"art"`
	IntId       int64      `xml:"int_id,attr"`
	Added       string     `xml:"added,attr"`
	Price       float64    `xml:"price,attr"`
	Cover       string     `xml:"cover,attr"`
	LastRelease string     `xml:"last_release,attr"`
	OnSale      int        `xml:"on_sale,attr"`
	FileId      int64      `xml:"file_id,attr"`
	Chars       int64      `xml:"chars,attr"`
	Type        int        `xml:"type,attr"`
	File        int64      `xml:"file,attr"`
	ShowPreview int        `xml:"show_preview,attr"`
	AllowRead   int        `xml:"allow_read,attr"`
	Lvl         int        `xml:"lvl,attr"`
	YearRating  YearRating `xml:"year_rating"`
	Authors     struct {
		Authors []Author `xml:"author"`
	} `xml:"authors"`
	ArtsRelations struct {
		ArtsRelations []ArtRelation `xml:"art_relation"`
	} `xml:"arts_relations"`
	TextDescription struct {
		Hidden Hidden `xml:"hidden"`
	} `xml:"text_description"`
}

type ArtRelation struct {
	XMLName  xml.Name `xml:"art_relation"`
	RelArt   int64    `xml:"rel_art,attr"`
	Relation int64    `xml:"relation,attr"`
}

type YearRating struct {
	XMLName  xml.Name `xml:"year_rating"`
	Quantity float64  `xml:"quantity,attr"`
}

type Author struct {
	XMLName          xml.Name `xml:"author"`
	Id               string   `xml:"id"`
	AuthorId         string   `xml:"id,attr"`
	SubjectId        int64    `xml:"subject_id"`
	Url              string   `xml:"url"`
	FirstName        string   `xml:"first-name"`
	MiddleName       string   `xml:"middle-name"`
	LastName         string   `xml:"last-name"`
	Email            string   `xml:"email"`
	FullNameGenitive string   `xml:"full-name-rodit"`
	Lvl              int      `xml:"lvl"`
	Relation         int      `xml:"relation"`
}

type Hidden struct {
	XMLName      xml.Name     `xml:"hidden"`
	TitleInfo    TitleInfo    `xml:"title-info"`
	DocumentInfo DocumentInfo `xml:"document-info"`
	PublishInfo  PublishInfo  `xml:"publish-info"`
}

type TitleInfo struct {
	XMLName    xml.Name `xml:"title-info"`
	Genres     []string `xml:"genre"`
	Author     Author   `xml:"author"`
	BookTitle  string   `xml:"book-title"`
	Annotation struct {
		P []string `xml:"p"`
	} `xml:"annotation"`
	Date      Date `xml:"date"`
	CoverPage struct {
		Image struct {
			XMLName xml.Name `xml:"image"`
			Href    string   `xml:"href,attr"`
		} `xml:"image"`
	} `xml:"coverpage"`
	Lang string `xml:"lang"`
}

type DocumentInfo struct {
	XMLName     xml.Name `xml:"document-info"`
	Id          string   `xml:"id"`
	Author      Author   `xml:"author"`
	ProgramUsed string   `xml:"program-used"`
	Date        Date     `xml:"date"`
	SrcUrls     []string `xml:"src-url"`
	SrcOcr      string   `xml:"src-ocr"`
}

type Date struct {
	XMLName xml.Name `xml:"date"`
	Value   string   `xml:"value,attr"`
}

type PublishInfo struct {
	XMLName   xml.Name `xml:"publish-info"`
	BookName  string   `xml:"book-name"`
	Publisher string   `xml:"publisher"`
	City      string   `xml:"city"`
	Year      int      `xml:"year"`
	ISBN      string   `xml:"isbn"`
	Sequence  Sequence `xml:"sequence"`
}

type Sequence struct {
	XMLName xml.Name `xml:"sequence"`
	Name    string   `xml:"name,attr"`
	Number  int      `xml:"number,attr"`
}
