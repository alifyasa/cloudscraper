package main

import (
	"log"
)

func AddScrapeTasks(scrapeTasks []ScrapeTask, scrapeTasksChannel chan<- ScrapeTask) {
	for _, scrapeTask := range scrapeTasks {
		go addScrapeTask(scrapeTask, scrapeTasksChannel)
	}
}

func addScrapeTask(scrapeTask ScrapeTask, scrapeTasksChannel chan<- ScrapeTask) {
	log.Printf("Adding scrape task %s\n", scrapeTask.url)
	scrapeTasksChannel <- scrapeTask
}
