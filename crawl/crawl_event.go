package crawl

import (
	"BaigeiCode/yearbook_api/db"
	"BaigeiCode/yearbook_api/model"
	"BaigeiCode/yearbook_api/vars"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlSudaEvents(year int64) error {
	err := db.DeleteAllEvents(year)
	if err != nil {
		return err
	}
	hasMore, idx, url, PageEvents := true, 0, vars.SUDA_EVENTS_INDEX_URL, make([]*model.Event, 0)
	for hasMore {
		PageEvents, hasMore, err = CrawlSudaEventsByPage(url, year)
		if err != nil {
			return err
		}
		_, err = db.MultiCreateEvents(PageEvents)
		if err != nil {
			return err
		}
		idx++
		url = fmt.Sprintf(vars.SUDA_EVENTS_INDEX_URL_FORMAT, idx)
	}
	return nil
}
func CrawlSudaEventsByPage(url string, nowYear int64) (events []*model.Event, hasMore bool, err error) {
	events = make([]*model.Event, 0)
	hasMore = true
	var rawHtml []byte
	var doc *goquery.Document
	rawHtml, err = GetHtmlByUrl(url)
	if err != nil {
		return nil, false, err
	}
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(string(rawHtml)))
	if err != nil {
		return nil, false, err
	}
	doc.Find("td.yx").Each(func(i int, ele *goquery.Selection) {
		if !hasMore {
			return
		}
		uri, _ := ele.Next().Next().Children().Attr("href")
		publishTime := ele.Text()
		year, _ := strconv.ParseInt(strings.Split(publishTime, "-")[0], 10, 64)
		if year == nowYear {
			events = append(events, &model.Event{
				Title:       ele.Next().Next().Text(),
				Uri:         uri,
				PublishTime: publishTime,
			})
		} else if year < nowYear {
			hasMore = false
			return
		}
	})
	return
}
func CrawlEventDetail(uri string) (string, error) {
	rawHtml, err := GetHtmlByUrl(fmt.Sprintf("%v%v", vars.SUDA_PLATFROM_PREFIX, uri))
	if err != nil {
		return "", err
	}
	content := ""
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(rawHtml)))
	if err != nil {
		return "", err
	}
	doc.Find("p").Each(func(i int, ele *goquery.Selection) {
		content += strings.Trim(ele.Text(), "\n")
	})
	return content, nil
}

// /suda_news/sdyw/202201/e47f5b64-b6d8-47d3-ac09-5aee4a5bc5ed.html
