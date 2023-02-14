package livelib_integration

import (
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	TitleProperty = "Название"
	// TypeProperty            = "Тип книги"
	AuthorsProperty    = "Авторы"
	PublishersProperty = "Издательство"
	// DownloadLinksProperty   = "Загрузки"
	DescriptionProperty    = "Описание"
	SequenceNameProperty   = "Название серии"
	SequenceNumberProperty = "Номер в серии"
	EditionSeriesProperty  = "Серия издания"
	GenresProperty         = "Жанры"
	// TagsProperty            = "Тэги"
	AwardsProperty = "Премии"
	// KnownIrrelevantProperty = "known-irrelevant-property"
)

func ReduceDetails(body *html.Node) (map[string][]string, error) {

	rdx := make(map[string][]string)

	bookHeader := match_node.NewEtc(atom.Section, "bc-header", true)
	if bh := match_node.Match(body, bookHeader); bh != nil {

		bookTitle := match_node.NewEtc(atom.H1, "bc__book-title ", true)
		if bt := match_node.Match(bh, bookTitle); bt != nil {
			rdx[TitleProperty] = []string{bt.FirstChild.Data}
		}

		bookAuthor := match_node.NewEtc(atom.A, "bc-author__link", true)
		for _, ba := range match_node.Matches(bh, bookAuthor, -1) {
			rdx[AuthorsProperty] = append(rdx[AuthorsProperty], ba.FirstChild.Data)
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
			sb := &strings.Builder{}
			if err := html.Render(sb, de); err != nil {
				return rdx, err
			}
			rdx[DescriptionProperty] = []string{sb.String()}
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
		for p, v := range reduceInfo(ie) {
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

		switch headingText {
		case "Серия:":
			rdx[EditionSeriesProperty] = values
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

func reduceInfo(node *html.Node) map[string][]string {
	rdx := make(map[string][]string)

	return rdx
}
