package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	scheme               = "https"
	LitResHost           = "litres.ru"
	wwwHost              = "www." + LitResHost
	partnersDownloadHost = "partnersdnld." + LitResHost

	coverPath        = "/pub/c/cover"
	cover330Path     = "/pub/c/cover_330"
	myBooksFreshPath = "/pages/my_books_fresh"
	detailedDataPath = "/static/ds/" + DetailedDataFilename
	biblioBookPath   = "/pages/biblio_book/"

	coverPathTemplate = "/{id}.jpg"
	pagePathTemplate  = "/page-{num}"

	DetailedDataFilename = "detailed_data.xml.gz"
)

func MyBooksFreshUrl(page int) *url.URL {

	pp := strings.Replace(pagePathTemplate, "{num}", strconv.Itoa(page), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   myBooksFreshPath + pp,
	}
}

func HrefUrl(href string) *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   href,
	}
}

func coverSizeUrl(id int64, coverSizePath string) *url.URL {
	cp := strings.Replace(coverPathTemplate, "{id}", strconv.FormatInt(id, 10), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   partnersDownloadHost,
		Path:   coverSizePath + cp,
	}
}

func CoverUrl(id int64) *url.URL {
	return coverSizeUrl(id, coverPath)
}

func Cover330Url(id int64) *url.URL {
	return coverSizeUrl(id, cover330Path)
}
