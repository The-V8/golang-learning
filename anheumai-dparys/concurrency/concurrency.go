package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	rch := make(chan result)

	for _, url := range urls {
		go func(u string) {
			rch <- result{u, wc(u)}
		}(url)
	}

	for n := 0; n < len(urls); n++ {
		result := <-rch
		results[result.string] = result.bool
	}

	return results
}
