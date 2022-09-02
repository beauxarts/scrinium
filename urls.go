package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	scheme               = "https"
	litresHost           = "litres.ru"
	wwwHost              = "www." + litresHost
	partnersDownloadHost = "partnersdnld." + litresHost

	coverPath        = "/pub/c/cover"
	myBooksFreshPath = "/pages/my_books_fresh"
	detailedDataPath = "/static/ds/detailed_data.xml.gz"

	coverPathTemplate = "/{id}.jpg"
	pagePathTemplate  = "/page-{num}"
)

func MyBooksFreshUrl(page int) *url.URL {

	pp := strings.Replace(pagePathTemplate, "{num}", strconv.Itoa(page), -1)

	u := &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   myBooksFreshPath + pp,
	}

	q := u.Query()

	q.Set("lite", "1")
	q.Set("gu_ajax", "true")

	u.RawQuery = q.Encode()

	return u
}

func CoverUrl(id int64) *url.URL {

	cp := strings.Replace(coverPathTemplate, "{id}", strconv.FormatInt(id, 10), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   partnersDownloadHost,
		Path:   coverPath + cp,
	}
}

func DetailedDataUrl() *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   detailedDataPath,
	}
}