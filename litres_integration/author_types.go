package litres_integration

import "golang.org/x/exp/maps"

type AuthorType int

const (
	AuthorUnknown AuthorType = iota
	AuthorDetails
	AuthorArts
)

var authorTypesStrings = map[AuthorType]string{
	AuthorUnknown: "author-unknown",
	AuthorDetails: "author-details",
	AuthorArts:    "author-similar",
}

func AllAuthorTypes() []AuthorType {
	return maps.Keys(authorTypesStrings)
}

func (at AuthorType) String() string {
	if ats, ok := authorTypesStrings[at]; ok {
		return ats
	}
	return authorTypesStrings[AuthorUnknown]
}

func ParseAuthorType(str string) AuthorType {
	for at, ats := range authorTypesStrings {
		if ats == str {
			return at
		}
	}
	return AuthorUnknown
}
