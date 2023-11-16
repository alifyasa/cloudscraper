package main

import "sync"

func MonitorScrapeTasks(
	scrapeTasksWaitGroup *sync.WaitGroup,
	monitorScrapeTasksWaitGroup *sync.WaitGroup,
	scrapeTasksChannel chan<- ScrapeTask,
	retryScrapeTasksChannel chan<- ScrapeTask,
) {
	defer monitorScrapeTasksWaitGroup.Done()
	scrapeTasksWaitGroup.Wait()
	close(retryScrapeTasksChannel)
	close(scrapeTasksChannel)
}
