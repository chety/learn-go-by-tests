package concurrency

type CheckWebSiteFunc func(string) bool
type URLCheck struct {
	url string
	ok  bool
}

func CheckWebSites(wc CheckWebSiteFunc, urls []string) map[string]bool {
	result := make(map[string]bool)
	resultChannel := make(chan URLCheck)

	for _, url := range urls {
		go func() {
			resultChannel <- URLCheck{url: url, ok: wc(url)}
		}()
	}
	for _, _ = range urls {
		r := <-resultChannel
		result[r.url] = r.ok
	}
	return result
}
