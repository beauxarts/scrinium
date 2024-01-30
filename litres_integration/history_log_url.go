package litres_integration

import (
	"net/url"
	"strconv"
	"strings"
)

func HistoryLogPageUrl(page int) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   LitResHost,
		Path:   strings.Replace(historyLogPathTemplate, "{0}", strconv.FormatInt(int64(page), 10), 1),
	}
}
