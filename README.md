# sitescraper
 
Scrape websites protected by Cloudflare using [`flaresolverr/flaresolverr`](https://github.com/FlareSolverr/FlareSolverr)

## How to Run

In `main.go`, fill the tasks with URLs. The program will use [`flaresolverr/flaresolverr`](https://github.com/FlareSolverr/FlareSolverr) to bypass Cloudflare protection and scrape the site.

```
go run ./cmd/.
```

The result will be written in `output/`.