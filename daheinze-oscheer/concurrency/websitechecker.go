package concurrency

// WebsiteChecker checks websites, what else do you think it should do?
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites What do you think will do this method?
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// Question to my reviewers: How does go recognize that all go func(...)
	// are done and accessable in the next lines?

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
