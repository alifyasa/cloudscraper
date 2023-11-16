package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

func ExecuteScrapeTasks(wg *sync.WaitGroup, scrapeTasksChannel <-chan ScrapeTask, retryScrapeTasksChannel chan<- ScrapeTask) {
	for scrapeTask := range scrapeTasksChannel {
		go executeScrapeTask(wg, scrapeTask, retryScrapeTasksChannel)
	}
}

func executeScrapeTask(wg *sync.WaitGroup, scrapetask ScrapeTask, retryScrapeTasksChannel chan<- ScrapeTask) {
	url := scrapetask.url
	requestBody := []byte(`{
		"cmd": "request.get",
		"url":"` + url + `",
		"maxTimeout": 120000
	}`)

	log.Printf("Executing scrape task %s\n", scrapetask.url)
	resp, err := http.Post("http://localhost:8191/v1", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		retryScrapeTasksChannel <- scrapetask
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		retryScrapeTasksChannel <- scrapetask
		return
	}

	handleScrapeTaskSuccess(wg, scrapetask, body)
	return
}

func handleScrapeTaskSuccess(wg *sync.WaitGroup, scrapeTask ScrapeTask, response []byte) {
	defer wg.Done()
	outputFolder := "output"
	// Check if the output folder exists, create it if not
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.MkdirAll(outputFolder, 0755)
		if err != nil {
			log.Println("Error creating output folder:", err)
			return
		}
	}

	filePath := fmt.Sprintf("%s/%d.txt", outputFolder, rand.Int()) // Path to the file inside the output folder

	err := os.WriteFile(filePath, response, 0644)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}

	log.Println("Data written to file successfully!")
}
