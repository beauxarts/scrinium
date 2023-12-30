package litres_integration

import "golang.org/x/exp/maps"

type SeriesType int

const (
	SeriesUnknown SeriesType = iota
	SeriesDetails
	SeriesArts
)

var seriesTypesStrings = map[SeriesType]string{
	SeriesUnknown: "series-unknown",
	SeriesDetails: "series-details",
	SeriesArts:    "series-similar",
}

func AllSeriesTypes() []SeriesType {
	return maps.Keys(seriesTypesStrings)
}

func (st SeriesType) String() string {
	if ats, ok := seriesTypesStrings[st]; ok {
		return ats
	}
	return seriesTypesStrings[SeriesUnknown]
}

func ParseSeriesType(str string) SeriesType {
	for st, sts := range seriesTypesStrings {
		if sts == str {
			return st
		}
	}
	return SeriesUnknown
}
