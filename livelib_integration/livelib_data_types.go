package livelib_integration

type LiveLibDataType int

const (
	LiveLibUnknownData LiveLibDataType = iota
	LiveLibDetails
)

var liveLibDataTypesString = map[LiveLibDataType]string{
	LiveLibUnknownData: "livelib-unknown-data",
	LiveLibDetails:     "livelib-details",
}

func AllLiveLibDataTypes() []LiveLibDataType {
	alldt := make([]LiveLibDataType, 0, len(liveLibDataTypesString)-1)
	for lldt := range liveLibDataTypesString {
		if lldt == LiveLibUnknownData {
			continue
		}
		alldt = append(alldt, lldt)
	}

	return alldt
}

func (lldt LiveLibDataType) String() string {
	if ats, ok := liveLibDataTypesString[lldt]; ok {
		return ats
	}
	return liveLibDataTypesString[LiveLibUnknownData]
}

func ParseLiveLibDataType(lldts string) LiveLibDataType {
	for lldt, str := range liveLibDataTypesString {
		if lldts == str {
			return lldt
		}
	}
	return LiveLibUnknownData
}
