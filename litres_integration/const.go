package litres_integration

const (
	httpsScheme          = "https"
	LitResHost           = "litres.ru"
	cvLitResHost         = "cv." + LitResHost
	apiLitResHost        = "api." + LitResHost
	wwwHost              = "www." + LitResHost
	partnersDownloadHost = "partnersdnld." + LitResHost

	coverPath    = "/pub/c/cover"
	cover330Path = "/pub/c/cover_330"
	cover415Path = "/pub/c/cover_415"

	foundationApiPath = "/foundation/api"

	artsDetailsPathTemplate     = foundationApiPath + "/arts/{id}"
	artsSimilarPathTemplate     = artsDetailsPathTemplate + "/similar"
	artsQuotesPathTemplate      = artsDetailsPathTemplate + "/quotes"
	artsUserReviewsPathTemplate = artsDetailsPathTemplate + "/user-reviews"
	artsFilesPathTemplate       = artsDetailsPathTemplate + "/files"
	artsReviewsPathTemplate     = artsDetailsPathTemplate + "/reviews"

	authorDetailsPathTemplate = foundationApiPath + "/authors/{id}"
	authorArtsPathTemplate    = authorDetailsPathTemplate + "/arts"

	seriesDetailsPathTemplate = foundationApiPath + "/series/{id}"
	seriesArtsPathTemplate    = seriesDetailsPathTemplate + "/arts"

	historyLogPathTemplate = "/pages/personal_cabinet_history_log/page-{0}"

	coverPathTemplate = "/{id}.jpg"
)
