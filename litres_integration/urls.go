package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	scheme               = "https"
	LitResHost           = "litres.ru"
	cvLitResHost         = "cv." + LitResHost
	wwwHost              = "www." + LitResHost
	partnersDownloadHost = "partnersdnld." + LitResHost

	coverPath        = "/pub/c/cover"
	cover330Path     = "/pub/c/cover_330"
	cover415Path     = "/pub/c/cover_415"
	myBooksFreshPath = "/pages/my_books_fresh"

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
