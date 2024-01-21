package litres_integration

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
