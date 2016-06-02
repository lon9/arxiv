package arxiv

import (
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	"path/filepath"
	"runtime"
	"strings"
)

// BaseURL is the base url of the arxiv.org.
const BaseURL = "https://arxiv.org/list"

// Crawler is crawler.
type Crawler struct {
	Genres []string
}

// NewCrawler is constructor of Crawler.
func NewCrawler() *Crawler {
	return &Crawler{}
}

// AddGenre adds genre.
func (c *Crawler) AddGenre(genre string) {
	c.Genres = append(c.Genres, genre)
}

// StartCrawl starts crawling.
func (c *Crawler) StartCrawl() (chan []Paper, chan error, chan bool) {
	ch := make(chan []Paper, runtime.NumCPU())
	errCh := make(chan error, runtime.NumCPU())
	doneCh := make(chan bool, 1)
	go c.crawl(ch, errCh, doneCh)
	return ch, errCh, doneCh
}

func (c *Crawler) crawl(ch chan []Paper, errCh chan error, doneCh chan bool) {
	for i := range c.Genres {
		url := BaseURL + "/" + c.Genres[i] + "/new"
		papers, err := c.scrape(url)
		if err != nil {
			errCh <- err
		} else {
			ch <- papers
		}
	}
	doneCh <- true
}

func (c *Crawler) scrape(url string) ([]Paper, error) {
	var papers []Paper
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	doc.Find("dl").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			s.Find("dt").Each(func(i int, s *goquery.Selection) {
				var paper Paper
				if id, exist := s.Find("span").Find("a").Attr("href"); exist {
					paper.ArxivKey = filepath.Base(id)
				}
				papers = append(papers, paper)
			})
			s.Find("dd").Each(func(i int, s *goquery.Selection) {
				papers[i].Title = strings.Replace(s.Find(".list-title").Text(), "Title: ", "", 1)
				s.Find(".list-authors").Find("a").Each(func(_ int, s *goquery.Selection) {
					var author Author
					author.Name = s.Text()
					//fmt.Println(author.Name)
					if href, exist := s.Attr("href"); exist {
						author.URL = href
					}
					papers[i].Authors = append(papers[i].Authors, author)
				})
				papers[i].Subject = s.Find(".list-subjects").Find(".primary-subject").Text()
				papers[i].Description = s.Find("p").Text()
			})
		}
	})
	return papers, nil
}
