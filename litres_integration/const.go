package litres_integration

const (
	httpsScheme = "https"

	LitResHost    = "litres.ru"
	cvLitResHost  = "cv." + LitResHost
	apiLitResHost = "api." + LitResHost
	wwwLitResHost = "www." + LitResHost

	foundationApiPath = "/foundation/api"

	artsDetailsPathTemplate = foundationApiPath + "/arts/{id}"
	artsSimilarPathTemplate = artsDetailsPathTemplate + "/similar"
	artsQuotesPathTemplate  = artsDetailsPathTemplate + "/quotes"
	artsFilesPathTemplate   = artsDetailsPathTemplate + "/files"
	artsReviewsPathTemplate = artsDetailsPathTemplate + "/reviews"

	authorDetailsPathTemplate = foundationApiPath + "/authors/{id}"
	authorArtsPathTemplate    = authorDetailsPathTemplate + "/arts"

	seriesDetailsPathTemplate = foundationApiPath + "/series/{id}"
	seriesArtsPathTemplate    = seriesDetailsPathTemplate + "/arts"

	usersMeArtsStatsPath = foundationApiPath + "/users/me/arts/stats"

	historyLogPathTemplate = "/pages/personal_cabinet_history_log/page-{0}"

	coverPathTemplate = "/pub/c/cover{size}/{id}.jpg"

	downloadPathTemplate = "/download_book/{id}/{file_id}/{filename}"
)
