package litres_integration

type ArtType int

const (
	ArtTypeText  ArtType = 0
	ArtTypeAudio ArtType = 1
	ArtTypePDF   ArtType = 4
)

var artTypeStrings = map[ArtType]string{
	ArtTypeText:  "Текст",
	ArtTypeAudio: "Аудио",
	ArtTypePDF:   "PDF",
}

func (at ArtType) String() string {
	return artTypeStrings[at]
}

func ParseArtType(ats string) ArtType {
	for at, str := range artTypeStrings {
		if str == ats {
			return at
		}
	}
	return ArtTypeText
}
