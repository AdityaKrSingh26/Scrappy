package linkedin

import (
	"strings"
	"time"

	"github.com/AdityaKrSingh26/Scrappy/core"
	"github.com/gocolly/colly/v2"
)

type LinkedinScraper struct{}

func (s *LinkedinScraper) Name() string {
	return "LinkedIn"
}

func (s *LinkedinScraper) Scrape() ([]core.Job, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		RandomDelay: 10 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	})

	var jobs []core.Job

	c.OnHTML(".jobs-search__results-list li", func(e *colly.HTMLElement) {
		title := e.ChildText(".base-search-card__title")
		company := e.ChildText(".base-search-card__subtitle")
		location := e.ChildText(".job-search-card__location")
		url := e.ChildAttr("a.base-card__full-link", "href")

		// Clean up data
		title = strings.TrimSpace(title)
		company = strings.TrimSpace(company)
		location = strings.TrimSpace(location)

		if title != "" && company != "" {
			job := core.Job{
				Title:    title,
				Company:  company,
				Location: location,
				URL:      url,
				// LinkedIn typically doesn't show description in list view
				Description: "See linked page for full description",
			}
			jobs = append(jobs, job)
		}
	})

	c.Visit("https://www.linkedin.com/jobs/search/?keywords=internship")

	c.Wait()
	return jobs, nil
}
