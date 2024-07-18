package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// go test -race
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement this sends the result values to the channel to avoid the race condition
			resultChannel <- result{u, wc(u)}
			// This causes a race condition
			// results[u] = wc(u)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
