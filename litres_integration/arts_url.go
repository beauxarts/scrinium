package litres_integration

import (
	"net/url"
	"strings"
)

var artsPathTemplates = map[ArtsType]string{
	ArtsDetails: artsDetailsPathTemplate,
	ArtsSimilar: artsSimilarPathTemplate,
	ArtsQuotes:  artsQuotesPathTemplate,
	ArtsFiles:   artsFilesPathTemplate,
	ArtsReviews: artsReviewsPathTemplate,
}

func ArtsUrl(at ArtsType, id string) *url.URL {
	pathTemplate := artsPathTemplates[at]
	if pathTemplate == "" {
		return nil
	}

	artsPath := strings.Replace(pathTemplate, "{id}", id, -1)

	return &url.URL{
		Scheme: scheme,
		Host:   apiLitResHost,
		Path:   artsPath,
	}
}
