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

	// is there a way to read everything from the channel without using len(urls)
	// there is... var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		// Receive expression this receives the result values from the channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
