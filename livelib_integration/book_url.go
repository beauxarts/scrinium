package livelib_integration

import (
	"net/url"
	"strings"
)

const (
	scheme      = "https"
	LiveLibHost = "livelib.ru"
	wwwHost     = "www." + LiveLibHost

	bookPath = "/book"

	bookIdTemplate = "/{id}"
)

func BookUrl(id string) *url.URL {

	bid := strings.Replace(bookIdTemplate, "{id}", id, -1)

	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   bookPath + bid,
	}
}
