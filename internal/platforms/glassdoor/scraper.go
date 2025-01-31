package glassdoor

import (
	"math/rand"
	"time"

	"github.com/AdityaKrSingh26/Scrappy/core"
	"github.com/gocolly/colly/v2"
)

type GlassdoorScraper struct{}

func (s *GlassdoorScraper) Name() string {
	return "Glassdoor"
}

func (s *GlassdoorScraper) Scrape() ([]core.Job, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.glassdoor.com"),
		colly.Async(true),
	)

	// Configure limits and delays
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomUserAgent())
	})

	var jobs []core.Job

	c.OnHTML(".react-job-listing", func(e *colly.HTMLElement) {
		job := core.Job{
			Title:       e.ChildText(".jobLink"),
			Company:     e.ChildText(".employerName"),
			Location:    e.ChildText(".location"),
			URL:         e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			Description: e.ChildText(".jobDescriptionContent"),
		}
		jobs = append(jobs, job)
	})

	// Handle pagination
	c.OnHTML("[data-test='pagination-next']", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.Visit("https://www.glassdoor.com/Internship/index.htm")

	c.Wait()
	return jobs, nil
}

func randomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
	}
	rand.Seed(time.Now().UnixNano())
	return userAgents[rand.Intn(len(userAgents))]
}
