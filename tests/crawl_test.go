package tests

import (
	"BaigeiCode/yearbook_api/db"
	"testing"
)

func TestCrawl(t *testing.T) {
	db.Init()
	// fmt.Println(crawl.CrawlSudaEvents(2022))
	// for _, e := range events {
	// 	s, _ := json.Marshal(e)
	// 	fmt.Println(string(s))
	// }
	// fmt.Println(crawl.CrawlEventDetail("/suda_news/sdyw/202201/e47f5b64-b6d8-47d3-ac09-5aee4a5bc5ed.html"))
}
