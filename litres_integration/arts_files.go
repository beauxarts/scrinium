package litres_integration

import (
	"net/url"
	"slices"
	"strconv"
	"strings"
)

//

var typeDescriptions = map[string]string{
	"fb2.zip":         "Читается всеми российскими электронными книгами и многими программами",
	"fb3":             "Развитие формата FB2",
	"txt":             "Можно открыть почти на любом устройстве",
	"txt.zip":         "Можно открыть почти на любом устройстве",
	"rtf.zip":         "Можно открыть почти на любом устройстве",
	"html":            "Можно читать книгу прямо в браузере",
	"html.zip":        "Можно читать книгу прямо в браузере",
	"mobi.prc":        "Подходит для электронных книг Amazon Kindle",
	"azw3":            "Подходит для электронных книг Amazon Kindle",
	"epub":            "Подходит для Apple Books и большинства приложений для чтения",
	"ios.epub":        "Подходит для Apple Books и большинства приложений для чтения",
	"audio/mpeg":      "Аудио книга в стандартном качестве",
	"application/zip": "Аудио книга в стандартном качестве",
	"audio/m4b":       "Версия для мобильного телефона",
	"a4.pdf":          "Открывается в программах Adobe Reader, Preview.app",
	"a6.pdf":          "Открывается в программах Adobe Reader, Preview.app",
	"17x24.pdf":       "Открывается в программах Adobe Reader, Preview.app",
	"application/pdf": "Открывается в программах Adobe Reader, Preview.app",
}

var preferredDownloadTypes = []string{
	"mobi.prc",
	"azw3",
	"epub",
	"audio/m4b",
	"application/pdf",
}

var preferredEncodingTypes = []string{
	"mobile_version_mp4",
	"zip_with_mp3",
	"pdf_book",
	"",
}

// additional_materials_pdf standard_quality_mp3_128kbps additional_materials_txt standard_quality_mp3  standard_quality_mp3_64kbps additional_materials_mp3

var neverDownloadEncodingTypes = []string{
	"introductory_fragment_pdf",
	"introductory_fragment_mp3",
}

type ArtsFilesData struct {
	ArtsId
	Filename     string  `json:"filename"`
	MIME         string  `json:"mime"`
	Extension    *string `json:"extension"`
	IsAdditional bool    `json:"is_additional"`
	ReleaseDate  *string `json:"release_date"`
	Pages        *int    `json:"pages"`
	Size         int     `json:"size"`
	Seconds      *int    `json:"seconds"`
	EncodingType string  `json:"encoding_type"`
}

type ArtsFiles struct {
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Payload struct {
		Data []ArtsFilesData `json:"data"`
	} `json:"payload"`
}

func (afd *ArtsFilesData) Type() string {
	if ext := afd.Extension; ext != nil {
		return *ext
	}
	return afd.MIME
}

func (afd *ArtsFilesData) TypeDescription() string {
	afdt := afd.Type()
	if td, ok := typeDescriptions[afdt]; ok {
		return td
	}
	return afdt
}

func (afd *ArtsFilesData) Url(id string) *url.URL {
	path := strings.Replace(downloadPathTemplate, "{id}", id, 1)
	path = strings.Replace(path, "{file_id}", strconv.FormatInt(int64(afd.Id), 10), 1)
	path = strings.Replace(path, "{filename}", afd.Filename, 1)

	if ext := afd.Extension; ext != nil {
		filenameWithExt := strings.Replace(afd.Filename, "zip", *ext, 1)
		path = strings.Replace(path, afd.Filename, filenameWithExt, 1)
	}

	return &url.URL{
		Scheme: httpsScheme,
		Host:   wwwLitResHost,
		Path:   path,
	}
}

func (af *ArtsFiles) DownloadsTypes() []*ArtsFilesData {
	afd := make([]*ArtsFilesData, 0)
	for _, d := range af.Payload.Data {
		if slices.Contains(preferredDownloadTypes, d.Type()) {
			if slices.Contains(neverDownloadEncodingTypes, d.EncodingType) {
				continue
			}
			if slices.Contains(preferredEncodingTypes, d.EncodingType) {
				afd = append(afd, &d)
			}
		}
	}
	// return all formats if data doesn't have preferred types
	if len(afd) == 0 {
		for _, d := range af.Payload.Data {
			afd = append(afd, &d)
		}
	}
	return afd
}
