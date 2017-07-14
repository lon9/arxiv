package arxiv

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	var c Crawler
	c.AddGenre("cs")
	ch, errCh, doneCh := c.StartCrawl()
L1:
	for {
		select {
		case papers := <-ch:
			t.Log(papers)
		case err := <-errCh:
			t.Fatal(err)
		case <-doneCh:
			close(ch)
			close(errCh)
			close(doneCh)
			break L1
		}
	}
}

func TestInitializer(t *testing.T) {
	c := NewCrawler([]string{"cs"})
	ch, errCh, doneCh := c.StartCrawl()
L1:
	for {
		select {
		case papers := <-ch:
			t.Log(papers)
		case err := <-errCh:
			t.Fatal(err)
		case <-doneCh:
			close(ch)
			close(errCh)
			close(doneCh)
			break L1
		}
	}
}
