package livelib_integration

import (
	"github.com/beauxarts/scrinium"
	"github.com/boggydigital/match_node"
	"github.com/boggydigital/nod"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	TitleProperty          = "Название"
	AuthorsProperty        = "Авторы"
	TranslatorsProperty    = "Переводчики"
	PublishersProperty     = "Издательство"
	YearPublishedProperty  = "Год издания"
	LanguageProperty       = "Язык"
	DescriptionProperty    = "Описание"
	SequenceNameProperty   = "Название серии"
	SequenceNumberProperty = "Номер в серии"
	GenresProperty         = "Жанры"
	AgeRatingProperty      = "Возрастные ограничения"
	TagsProperty           = "Теги"
	ISBNProperty           = "ISBN"
	ImageProperty          = "Обложка"
)

func Reduce(body *html.Node) (map[string][]string, error) {

	rdx := make(map[string][]string)

	bookHeader := match_node.NewEtc(atom.Section, "bc-header", true)
	if bh := match_node.Match(body, bookHeader); bh != nil {

		bookTitle := match_node.NewEtc(atom.H1, "bc__book-title ", true)
		if bt := match_node.Match(bh, bookTitle); bt != nil {
			rdx[TitleProperty] = []string{bt.FirstChild.Data}
		}

		bookAuthor := match_node.NewEtc(atom.A, "bc-author__link", false)
		for _, ba := range match_node.Matches(bh, bookAuthor, -1) {
			rdx[AuthorsProperty] = append(rdx[AuthorsProperty], ba.FirstChild.Data)
		}
	}

	bookImage := match_node.NewId(atom.Img, "main-image-book")
	if bi := match_node.Match(body, bookImage); bi != nil {
		if src := scrinium.GetAttribute(bi, "src"); src != "" {
			rdx[ImageProperty] = []string{src}
		}
	}

	genresList := match_node.NewEtc(atom.Ul, "bc-genre__list", true)
	if gl := match_node.Match(body, genresList); gl != nil {
		for _, a := range match_node.Matches(gl, match_node.NewEtc(atom.A, "", false), -1) {
			rdx[GenresProperty] = append(rdx[GenresProperty], a.FirstChild.Data)
		}
	}

	annotation := match_node.NewEtc(atom.Div, "bc-annotate", false)
	if an := match_node.Match(body, annotation); an != nil {

		desc := match_node.NewId(atom.Div, "lenta-card__text-edition-escaped")
		if de := match_node.Match(an, desc); de != nil {
			rdx[DescriptionProperty] = []string{
				trimNewLinesWhitespace(match_node.TextContent(de))}
		}
	}

	edition := match_node.NewEtc(atom.Table, "bc-edition", true)
	if en := match_node.Match(body, edition); en != nil {
		for p, v := range reduceEdition(en) {
			rdx[p] = v
		}
	}

	info := match_node.NewEtc(atom.Div, "bc-info", true)
	if ie := match_node.Match(body, info); ie != nil {
		infoRdx, err := reduceInfo(ie)
		if err != nil {
			return rdx, err
		}
		for p, v := range infoRdx {
			rdx[p] = v
		}
	}

	return rdx, nil
}

func reduceEdition(node *html.Node) map[string][]string {
	rdx := make(map[string][]string)
	heading := match_node.NewEtc(atom.Td, "bc-add-info__heading", false)
	for _, htd := range match_node.Matches(node, heading, -1) {
		headingText := htd.FirstChild.Data
		values := make([]string, 0)

		if headingSibling := htd.NextSibling.NextSibling; headingSibling != nil {
			link := match_node.NewEtc(atom.A, "bc-edition__link", true)
			for _, ae := range match_node.Matches(headingSibling, link, -1) {
				values = append(values, strings.TrimSpace(ae.FirstChild.Data))
			}
		}

		if len(values) == 0 {
			for sibling := htd.NextSibling; sibling != nil; sibling = sibling.NextSibling {
				if sibling.DataAtom != atom.Td {
					continue
				}
				values = append(values, sibling.FirstChild.Data)
			}
		}

		switch headingText {
		case "Серия:":
			fallthrough
		case "Цикл:":
			for _, v := range values {
				if st, sn, ok := strings.Cut(v, ","); ok {
					rdx[SequenceNameProperty] = append(rdx[SequenceNameProperty], st)
					rdx[SequenceNumberProperty] = append(rdx[SequenceNumberProperty], strings.TrimSpace(sn))
				}
			}
		case "Издательство:":
			rdx[PublishersProperty] = values
		}
	}
	return rdx
}

var liveLibProperties = map[string]string{
	"Перевод":             TranslatorsProperty,
	ISBNProperty:          ISBNProperty,
	YearPublishedProperty: YearPublishedProperty,
	LanguageProperty:      LanguageProperty,
	AgeRatingProperty:     AgeRatingProperty,
	GenresProperty:        GenresProperty,
	TagsProperty:          TagsProperty,
}

func reduceInfo(node *html.Node) (map[string][]string, error) {
	rdx := make(map[string][]string)

	pars := match_node.NewAtom(atom.P)
	for _, p := range match_node.Matches(node, pars, -1) {

		innerText := match_node.TextContent(p)
		innerText = strings.TrimSpace(innerText)

		if p, v, ok := strings.Cut(innerText, ":"); ok {
			if property, ok := liveLibProperties[p]; ok {
				rdx[property] = trimValues(v)
			} else {
				nod.Log("%s: %s", p, v)
			}
		}
	}

	return rdx, nil
}

func trimNewLinesWhitespace(s string) string {
	return strings.TrimSpace(strings.Replace(s, "\n", "", -1))
}

func trimValues(v string) []string {
	values := strings.Split(v, ",")
	tv := make([]string, len(values))
	for i, v := range values {
		tv[i] = strings.TrimSpace(v)
	}
	return tv
}
