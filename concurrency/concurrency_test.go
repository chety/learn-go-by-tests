package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func checkWebSiteMock(url string) bool {
	return url != "https://google.com"
}

func TestCheckWebSites(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://duckduckgo.com",
	}

	want := map[string]bool{
		"https://google.com":     false,
		"https://yahoo.com":      true,
		"https://duckduckgo.com": true,
	}

	t.Run("Should return the correct ok", func(t *testing.T) {
		got := CheckWebSites(checkWebSiteMock, urls)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("CheckWebSites: want %v, got %v", want, got)
		}
	})

}

func slowMockCheckWebSites(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebSites(b *testing.B) {
	urls := make([]string, 100)
	for i, _ := range urls {
		urls[i] = "https://google.com"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebSites(slowMockCheckWebSites, urls)
	}
}
