package converter

import (
	"chproxy/app/model"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ConvertResponse(html io.Reader) (dat io.Reader) {
	doc, _ := goquery.NewDocumentFromReader(html)

	title := func() string {
		return doc.Find("body > div > div.title").Text()
	}()

	l := make([]model.Res, 0)
	doc.Find("body > div > div.thread > div.post").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div.meta > span.name").Html()

		email := func() string {
			s, _ := s.Find("div.meta > span.name > b > a").Attr("href")
			return strings.TrimPrefix(s, "mailto:")
		}()

		date := s.Find("div.meta > span.date").Text()

		id := strings.TrimPrefix(
			s.Find("div.meta > span.uid").Text(),
			"ID:",
		)

		text := func() string {
			t, _ := s.Find("div.message > span.escaped").Html()

			lines := make([]string, 0)
			for _, l := range strings.Split(t, "<br/>") {
				l = strings.TrimLeft(l, " ")
				l = strings.TrimRight(l, " ")

				lines = append(lines, l)
			}

			return strings.Join(lines, "\n")
		}()

		l = append(
			l,
			model.Res{
				Name:  name,
				EMail: email,
				Date:  date,
				ID:    id,
				Text:  text,
			},
		)
	})

	t := model.Thread{
		Title:   title,
		ResList: l,
	}

	return strings.NewReader(t.ToDat())
}
