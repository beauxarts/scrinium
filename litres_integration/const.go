package litres_integration

const (
	scheme               = "https"
	LitResHost           = "litres.ru"
	cvLitResHost         = "cv." + LitResHost
	apiLitResHost        = "api." + LitResHost
	wwwHost              = "www." + LitResHost
	partnersDownloadHost = "partnersdnld." + LitResHost

	coverPath                    = "/pub/c/cover"
	cover330Path                 = "/pub/c/cover_330"
	cover415Path                 = "/pub/c/cover_415"
	myBooksFreshPath             = "/pages/my_books_fresh"
	myBooksFreshPagePathTemplate = "/page-{num}"

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

	coverPathTemplate = "/{id}.jpg"
)
