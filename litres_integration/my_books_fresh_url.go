package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

func MyBooksFreshUrl(page int) *url.URL {

	pp := strings.Replace(myBooksFreshPagePathTemplate, "{num}", strconv.Itoa(page), -1)

	return &url.URL{
		Scheme: scheme,
		Host:   wwwHost,
		Path:   myBooksFreshPath + pp,
	}
}
