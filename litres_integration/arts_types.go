package litres_integration

import "golang.org/x/exp/maps"

type ArtsType int

const (
	ArtsUnknown ArtsType = iota
	ArtsDetails
	ArtsSimilar
	ArtsQuotes
	ArtsUserReviews
	ArtsFiles
	ArtsReviews
)

var artsTypesStrings = map[ArtsType]string{
	ArtsUnknown:     "arts-unknown",
	ArtsDetails:     "arts-details",
	ArtsSimilar:     "arts-similar",
	ArtsQuotes:      "arts-quotes",
	ArtsUserReviews: "arts-user-reviews",
	ArtsFiles:       "arts-files",
	ArtsReviews:     "arts-reviews",
}

func AllArtsTypes() []ArtsType {
	return maps.Keys(artsTypesStrings)
}

func (at ArtsType) String() string {
	if ats, ok := artsTypesStrings[at]; ok {
		return ats
	}
	return artsTypesStrings[ArtsUnknown]
}

func ParseArtsType(str string) ArtsType {
	for at, ats := range artsTypesStrings {
		if ats == str {
			return at
		}
	}
	return ArtsUnknown
}
