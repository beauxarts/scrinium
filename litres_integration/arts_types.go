package litres_integration

type ArtsType int

const (
	ArtsUnknown ArtsType = iota
	ArtsDetails
	ArtsSimilar
	ArtsQuotes
	ArtsFiles
	ArtsReviews
)

var artsTypesStrings = map[ArtsType]string{
	ArtsUnknown: "arts-unknown",
	ArtsDetails: "arts-details",
	ArtsSimilar: "arts-similar",
	ArtsQuotes:  "arts-quotes",
	ArtsFiles:   "arts-files",
	ArtsReviews: "arts-reviews",
}

func AllArtsTypes() []ArtsType {
	aat := make([]ArtsType, 0, len(artsTypesStrings)-1)
	for at := range artsTypesStrings {
		if at == ArtsUnknown {
			continue
		}
		aat = append(aat, at)
	}

	return aat
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
