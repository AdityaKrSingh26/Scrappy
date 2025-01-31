package wellfound

import (
	"github.com/AdityaKrSingh26/Scrappy/core"
	"github.com/gocolly/colly/v2"
)

type WellfoundScraper struct{}

func (s *WellfoundScraper) Name() string {
	return "Wellfound"
}

func (s *WellfoundScraper) Scrape() ([]core.Job, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("wellfound.com"),
	)

	var jobs []core.Job

	c.OnHTML(".internship-listing", func(e *colly.HTMLElement) {
		job := core.Job{
			Title:       e.ChildText("h2.title"),
			Company:     e.ChildText(".company-name"),
			Location:    e.ChildText(".location"),
			URL:         e.Request.AbsoluteURL(e.Attr("href")),
			Description: e.ChildText(".description"),
		}
		jobs = append(jobs, job)
	})

	c.Visit("https://wellfound.com/jobs")

	return jobs, nil
}
