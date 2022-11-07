package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var sites = []string{
	"http://www.nytimes.com/services/xml/rss/nyt/International.xml",
	"http://www.nytimes.com/services/xml/rss/nyt/Technology.xml",
	"http://feeds.bbci.co.uk/news/world/rss.xml",
	"http://www.aljazeera.com/xml/rss/all.xml",
	"http://defence-blog.com/feed",
	"http://www.e-ir.info/category/blogs/feed",
	"http://www.globalissues.org/news/feed",
	"http://rss.cnn.com/rss/edition_world.rss",
	"http://timesofindia.indiatimes.com/rssfeeds/296589292.cms",
	"http://feeds.washingtonpost.com/rss/world",
	"https://feeds.a.dj.com/rss/RSSOpinion.xml",
	"https://feeds.a.dj.com/rss/RSSWorldNews.xml",
	"https://feeds.a.dj.com/rss/WSJcomUSBusiness.xml",
	"https://feeds.a.dj.com/rss/WSJcomUSBusiness.xml",
	"https://feeds.a.dj.com/rss/RSSWSJD.xml",
	"http://xml2.corriereobjects.it/rss/english.xml",
}

var debug bool

// Feed minimal data.
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title string `xml:"title"` // BBC News - World
		Item  []struct {
			Title   string `xml:"title"`   // Thai king strips consort ...
			PubDate string `xml:"pubDate"` // Mon, 21 Oct 2019 17:00:48...
		} `xml:"item"`
	} `xml:"channel"`
}

type result struct {
	headline string
	err      error
}

type search func(context.Context) result

func first(ctx context.Context, searchers ...search) result {
	c := make(chan result, len(searchers))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// For each searcher,
	for _, s := range searchers {
		go func(s search) {
			c <- s(ctx)
		}(s)
	}

	for {
		select {
		case <-ctx.Done():
			return result{err: ctx.Err()}
		case r := <-c:
			if r.err == nil && len(r.headline) > 0 {
				return r
			}
			log.Printf("empty headline: %v", r.err)
		}
	}
}

func randomHeadline(r io.Reader) (string, error) {
	dec := xml.NewDecoder(r)
	var feed Feed
	if err := dec.Decode(&feed); err != nil {
		return "", err
	}
	size := len(feed.Channel.Item)
	randomItem := feed.Channel.Item[rand.Intn(size)]
	// v := fmt.Sprintf("%s %s [%s]", randomItem.Title, randomItem.PubDate, feed.Channel.Title)
	v := fmt.Sprintf("%s [%s]", randomItem.Title, feed.Channel.Title)
	return v, nil
}

func createSearcher(link string) search {
	f := func(ctx context.Context) result {
		if debug {
			log.Printf("lookup %s", link)
		}
		resultStream := make(chan result)

		go func() {
			req, err := http.NewRequestWithContext(ctx, "GET", link, nil)
			if err != nil {
				resultStream <- result{err: err}
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				resultStream <- result{err: err}
				return
			}
			defer resp.Body.Close()
			headline, err := randomHeadline(resp.Body)
			resultStream <- result{headline, err}
		}()

		select {
		case v := <-resultStream:
			return v
		case <-ctx.Done():
			return result{err: ctx.Err()}
		}

	}
	return f
}

func main() {

	// Use a subset of feeds.
	rand.Seed(time.Now().UnixNano())
	sample := append([]string(nil), sites...)
	rand.Shuffle(len(sites), func(i, j int) { sample[i], sample[j] = sample[j], sample[i] })
	sites = sample[:8] // Beware panic.

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	var searchers []search
	for _, link := range sites {
		searchers = append(searchers, createSearcher(link))
	}

	result := first(ctx, searchers...)

	if result.err != nil || result.headline == "" {
		fmt.Printf("No headline available at this time.\n")
		os.Exit(1)
	}
	fmt.Println(result.headline)

}
