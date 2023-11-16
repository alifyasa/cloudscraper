package main

import (
	"log"
	"sync"
)

type ScrapeTask struct {
	url        string
	retryCount int
}

func main() {

	var scrapeTasksWaitGroup sync.WaitGroup
	var monitorScrapeTasksWaitGroup sync.WaitGroup

	scrapeTasksChannel := make(chan ScrapeTask)
	retryScrapeTasksChannel := make(chan ScrapeTask)

	tasks := []ScrapeTask{
		{url: "https://example.com/", retryCount: 0},
		{url: "https://example.com/", retryCount: 0},
		{url: "https://example.com/", retryCount: 0},
	}

	log.Println("YOOOOO")

	scrapeTasksWaitGroup.Add(len(tasks))
	go AddScrapeTasks(tasks, scrapeTasksChannel)
	go ExecuteScrapeTasks(&scrapeTasksWaitGroup, scrapeTasksChannel, retryScrapeTasksChannel)
	go RetryScrapeTasks(retryScrapeTasksChannel, scrapeTasksChannel)
	monitorScrapeTasksWaitGroup.Add(1)
	go MonitorScrapeTasks(
		&scrapeTasksWaitGroup,
		&monitorScrapeTasksWaitGroup,
		scrapeTasksChannel,
		retryScrapeTasksChannel,
	)

	monitorScrapeTasksWaitGroup.Wait()
}
