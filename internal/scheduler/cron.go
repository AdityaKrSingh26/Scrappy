package scheduler

import (
	"fmt"
	"time"

	"github.com/AdityaKrSingh26/Scrappy/cmd"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// Scheduler is responsible for managing cron jobs
type Scheduler struct {
	cron *cron.Cron
}

// NewScheduler creates a new Scheduler instance with a cron job manager
func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(cron.WithSeconds()), // Use seconds-level precision
	}
}

// Start begins the cron scheduler
func (s *Scheduler) Start() {
	logrus.Info("Starting the cron scheduler...")

	// Add a cron job to run the scraping task every day at 9:00 AM
	_, err := s.cron.AddFunc("0 0 9 * * *", func() {
		logrus.Info("Running scheduled scraping task...")
		err := cmd.RunScrapeCommand() // Call your scrape function from scrape.go
		if err != nil {
			logrus.Errorf("Error running scrape command: %v", err)
		} else {
			logrus.Info("Scraping task completed successfully.")
		}
	})

	if err != nil {
		logrus.Fatalf("Error scheduling scraping task: %v", err)
	}

	// Start the cron scheduler
	s.cron.Start()

	// Keep the scheduler running
	select {}
}

// Stop gracefully stops the cron scheduler
func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	logrus.Info("Cron scheduler stopped.")
	select {
	case <-ctx.Done():
		logrus.Info("All cron jobs have completed.")
	}
}

// ManualScrape allows the user to run a manual scraping task
func (s *Scheduler) ManualScrape() {
	logrus.Info("Running manual scraping task...")
	err := cmd.RunScrapeCommand() // Call the scrape command manually
	if err != nil {
		logrus.Errorf("Error running manual scrape command: %v", err)
	} else {
		logrus.Info("Manual scraping task completed successfully.")
	}
}

// LogNextScheduledRun prints out the time of the next scheduled scraping task
func (s *Scheduler) LogNextScheduledRun() {
	entries := s.cron.Entries()
	if len(entries) > 0 {
		nextRun := entries[0].Next
		fmt.Printf("Next scheduled scraping task will run at: %v\n", nextRun.Format(time.RFC1123))
	} else {
		fmt.Println("No scheduled tasks found.")
	}
}
