package litres_integration

import (
	"math"
	"net/url"
	"strconv"
	"strings"
)

type CoverSize int

const (
	Size330 CoverSize = 330
	Size415 CoverSize = 415
	SizeMax CoverSize = math.MaxInt
)

func coverSizeUrl(id int64, host string, coverSizePath string) *url.URL {
	cp := strings.Replace(coverPathTemplate, "{id}", strconv.FormatInt(id, 10), -1)

	return &url.URL{
		Scheme: httpsScheme,
		Host:   host,
		Path:   coverSizePath + cp,
	}
}

func pathSize(size CoverSize) string {
	path := coverPath
	switch size {
	case Size330:
		path = cover330Path
	case Size415:
		path = cover415Path
	case SizeMax:
		// do nothing - that's the default value
	}

	return path
}

func CoverUrl(id int64, size CoverSize) *url.URL {
	return coverSizeUrl(id, cvLitResHost, pathSize(size))
}

func PartnerCoverUrl(id int64, size CoverSize) *url.URL {
	return coverSizeUrl(id, partnersDownloadHost, pathSize(size))
}
