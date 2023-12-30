package litres_integration

type LitResDataType int

const (
	LitResUnknownData LitResDataType = iota
	LitResMyBooksFresh
	LitResMyBooksDetails
)

var litresDataTypeStrings = map[LitResDataType]string{
	LitResUnknownData:    "litres-unknown-data",
	LitResMyBooksFresh:   "litres-my-books-fresh",
	LitResMyBooksDetails: "litres-my-books-details",
}

func AllLitResDataTypes() []LitResDataType {
	alrdt := make([]LitResDataType, 0, len(litresDataTypeStrings)-1)
	for lrdt := range litresDataTypeStrings {
		if lrdt == LitResUnknownData {
			continue
		}
		alrdt = append(alrdt, lrdt)
	}

	return alrdt
}

func (lrdt LitResDataType) String() string {
	if dts, ok := litresDataTypeStrings[lrdt]; ok {
		return dts
	}
	return ""
}

func ParseLitResDataType(lrdts string) LitResDataType {
	for dt, str := range litresDataTypeStrings {
		if str == lrdts {
			return dt
		}
	}
	return LitResUnknownData
}
