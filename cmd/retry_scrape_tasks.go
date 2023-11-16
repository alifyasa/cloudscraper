package main

import "log"

func RetryScrapeTasks(retryScrapeTasksChannel <-chan ScrapeTask, scrapeTasksChannel chan<- ScrapeTask) {
	for scrapeTask := range retryScrapeTasksChannel {
		scrapeTask.retryCount += 1
		log.Printf("Retry number %d for scrape task %s\n", scrapeTask.retryCount, scrapeTask.url)
		scrapeTasksChannel <- scrapeTask
	}
}
